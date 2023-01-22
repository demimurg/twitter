package postgres

import (
	"context"
	"database/sql"
    "errors"
    "fmt"
    "github.com/jackc/pgx/v5/pgconn"
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

func (u *userRepo) Add(ctx context.Context, name, email, caption string, birthDate time.Time) (id int, err error) {
	row := u.db.QueryRowContext(ctx, `
        INSERT INTO users (full_name, email, caption, birth_date)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `, name, email, caption, birthDate)

	if err := row.Scan(&id); err != nil {
        var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) && pgErr.Code == "23505" {
            // violates unique constraint "users_email_key"
            return 0, usecase.ErrUserExists
        }
		return 0, fmt.Errorf("insert user to db: %w", err)
	}
	return id, nil
}

func (u *userRepo) Get(ctx context.Context, userID int) (*entity.User, error) {
	row := u.db.QueryRowContext(ctx, `
        SELECT id, full_name, email, caption, birth_date, deleted_at
        FROM users
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
        FROM users
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
        FROM users
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
        UPDATE users
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
        UPDATE users
        SET deleted_at = now()
        WHERE id = $2 AND deleted_at IS NULL
    `, userID)
	if err != nil {
		return fmt.Errorf("delete user from db: %w", err)
	}
	return nil
}
