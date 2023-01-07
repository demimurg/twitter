package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/demimurg/twitter/internal/entity"
	"github.com/demimurg/twitter/internal/usecase"
)

func NewUserRepository(db *sql.DB) usecase.UserRepository {
	return &userRepo{db: db}
}

type userRepo struct {
	db *sql.DB
}

func (u *userRepo) Add(ctx context.Context, name, email, caption string, birthDate time.Time) (int, error) {
	row := u.db.QueryRowContext(ctx, `
        INSERT INTO "user" (full_name, email, caption, birth_date, created_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `, name, email, caption, birthDate, time.Now())

	var userID int
	if err := row.Scan(&userID); err != nil {
		return 0, fmt.Errorf("insert user to db: %w", err)
	}
	return userID, nil
}

func (u *userRepo) Get(ctx context.Context, userID int) (*entity.User, error) {
	row := u.db.QueryRowContext(ctx, `
        SELECT id, full_name, email, caption, birth_date, deleted_at
        FROM "user"
        WHERE id = $1 AND deleted_at IS NULL
    `, userID)

	var user entity.User
	err := row.Scan(
		&user.ID, &user.FullName, &user.Email,
		&user.Caption, &user.BirthDate, &user.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("select user from db: %w", err)
	}
	return &user, nil
}

func (u *userRepo) GetAll(ctx context.Context, limit int) ([]entity.User, error) {
	rows, err := u.db.QueryContext(ctx, `
        SELECT id, full_name, email, caption, birth_date, deleted_at
        FROM "user"
        WHERE deleted_at IS NULL
        LIMIT $1
    `, limit)
	if err != nil {
		return nil, fmt.Errorf("select users from db: %w", err)
	}
	defer rows.Close()

	var (
		user  entity.User
		users []entity.User
	)
	for rows.Next() {
		err := rows.Scan(
			&user.ID, &user.FullName, &user.Email,
			&user.Caption, &user.BirthDate, &user.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan user from db: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}
func (u *userRepo) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	row := u.db.QueryRowContext(ctx, `
        SELECT id, full_name, email, caption, birth_date, deleted_at
        FROM "user"
        WHERE email = $1 AND deleted_at IS NULL
    `, email)

	var user entity.User
	err := row.Scan(
		&user.ID, &user.FullName, &user.Email,
		&user.Caption, &user.BirthDate, &user.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("select user from db: %w", err)
	}
	return &user, nil
}

func (u *userRepo) UpdateCaption(ctx context.Context, userID int, caption string) error {
	_, err := u.db.ExecContext(ctx, `
        UPDATE "user"
        SET caption = $1
        WHERE id = $2 AND deleted_at IS NULL
    `, caption, userID)
	if err != nil {
		return fmt.Errorf("update user caption in db: %w", err)
	}
	return nil
}

func (u *userRepo) Delete(ctx context.Context, userID int) error {
	_, err := u.db.ExecContext(ctx, `
        UPDATE "user"
        SET deleted_at = $1
        WHERE id = $2 AND deleted_at IS NULL
    `, time.Now(), userID)
	if err != nil {
		return fmt.Errorf("delete user from db: %w", err)
	}
	return nil
}
