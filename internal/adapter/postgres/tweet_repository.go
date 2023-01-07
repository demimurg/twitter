package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/demimurg/twitter/internal/entity"
	"github.com/demimurg/twitter/internal/usecase"
)

func NewTweetRepository(db *sql.DB) usecase.TweetRepository {
	return &tweetRepo{db: db}
}

type tweetRepo struct {
	db *sql.DB
}

func (t *tweetRepo) Add(ctx context.Context, userID int, tweetText string) error {
	_, err := t.db.ExecContext(ctx, `
        INSERT INTO tweet (user_id, text)
        VALUES ($1, $2)
    `, userID, tweetText)
	if err != nil {
		return fmt.Errorf("insert tweet to db: %w", err)
	}
	return nil
}

func (t *tweetRepo) Update(ctx context.Context, tweetID int, newText string) error {
	_, err := t.db.ExecContext(ctx, `
        UPDATE tweet
        SET text = $1
        WHERE id = $2
    `, newText, tweetID)
	if err != nil {
		return fmt.Errorf("update tweet in db: %w", err)
	}
	return nil
}

func (t *tweetRepo) AddComment(ctx context.Context, userID, tweetID int, commentText string) error {
	_, err := t.db.ExecContext(ctx, `
        INSERT INTO comment (user_id, tweet_id, text)
        VALUES ($1, $2, $3)
    `, userID, tweetID, commentText)
	if err != nil {
		return fmt.Errorf("insert comment to db: %w", err)
	}
	return nil
}

func (t *tweetRepo) UpdateComment(ctx context.Context, commentID int, newText string) error {
	_, err := t.db.ExecContext(ctx, `
        UPDATE comment
        SET text = $1
        WHERE id = $2
    `, newText, commentID)
	if err != nil {
		return fmt.Errorf("update comment in db: %w", err)
	}
	return nil
}

func (t *tweetRepo) GetLatest(ctx context.Context, userID int, limit int) ([]entity.Tweet, error) {
	rows, err := t.db.QueryContext(ctx, `
        SELECT id, user_id, text, likes, created_at
        FROM tweet
        WHERE user_id = $1 AND deleted_at IS NULL
        ORDER BY created_at DESC
        LIMIT $2
    `, userID, limit)
	if err != nil {
		return nil, fmt.Errorf("select tweet tweets from db: %w", err)
	}
	defer rows.Close()

	var (
		tweet  entity.Tweet
		tweets []entity.Tweet
	)
	for rows.Next() {
		err := rows.Scan(
			&tweet.ID, &tweet.UserID, &tweet.Text,
			&tweet.Likes, &tweet.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan tweet from db: %w", err)
		}
		tweets = append(tweets, tweet)
	}

	return tweets, nil
}
