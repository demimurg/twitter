package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/demimurg/twitter/internal/usecase"
)

func NewFollowerRepository(db *sql.DB) usecase.FollowerRepository {
	return &followerRepo{db: db}
}

type followerRepo struct {
	db *sql.DB
}

func (f *followerRepo) Add(ctx context.Context, followerID, toUserID int) error {
	_, err := f.db.ExecContext(ctx, `
        INSERT INTO follower (follower_id, followee_id)
        VALUES ($1, $2)
    `, followerID, toUserID)
	if err != nil {
		return fmt.Errorf("insert follower to db: %w", err)
	}
	return nil
}

func (f *followerRepo) Remove(ctx context.Context, followerID, fromUserID int) error {
	_, err := f.db.ExecContext(ctx, `
        DELETE FROM follower
        WHERE follower_id = $1 AND followee_id = $2
            AND deleted_at IS NULL
    `, followerID, fromUserID)
	if err != nil {
		return fmt.Errorf("delete follower from db: %w", err)
	}
	return nil
}

func (f *followerRepo) GetFollowing(ctx context.Context, userID, topN int) ([]int, error) {
	rows, err := f.db.QueryContext(ctx, `
        SELECT followee_id FROM follower
        WHERE follower_id = $1 AND deleted_at IS NULL
        LIMIT $2
    `, userID, topN)
	if err != nil {
		return nil, fmt.Errorf("select following users ids: %w", err)
	}
	defer rows.Close()

	var (
		id       int
		usersIDs []int
	)
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan following user id: %w", err)
		}
		usersIDs = append(usersIDs, id)
	}

	return usersIDs, nil
}

func (f *followerRepo) GetFollowers(ctx context.Context, userID, topN int) ([]int, error) {
	rows, err := f.db.QueryContext(ctx, `
        SELECT follower_id FROM follower
        WHERE followee_id = $1 AND deleted_at IS NULL
        LIMIT $2
    `, userID, topN)
	if err != nil {
		return nil, fmt.Errorf("select followers ids: %w", err)
	}
	defer rows.Close()

	var (
		id       int
		usersIDs []int
	)
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan follower id: %w", err)
		}
		usersIDs = append(usersIDs, id)
	}

	return usersIDs, nil
}
