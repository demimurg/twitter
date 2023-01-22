package inmem

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/demimurg/twitter/internal/entity"
	"github.com/demimurg/twitter/internal/usecase"
)

var errNoSuchUser = errors.New("there is no such user in memory")

func NewUserRepository() usecase.UserRepository {
	return &userRepo{users: make([]entity.User, 0, 100)} // initial tweets size 100 users
}

type userRepo struct {
	users []entity.User
}

func (u *userRepo) Add(_ context.Context, name, email, caption string, birthDate time.Time) (int, error) {
    for _, user := range u.users {
        if user.Email == email {
            return 0, usecase.ErrUserExists
        }
    }

	id := len(u.users)
	u.users = append(u.users, entity.User{
		ID:        id,
		Email:     email,
		FullName:  name,
		Caption:   caption,
		BirthDate: birthDate,
	})
	return id, nil
}

func (u *userRepo) Get(_ context.Context, userID int) (*entity.User, error) {
	if userID >= len(u.users) || userID < 0 {
		return nil, errNoSuchUser
	}

	return &u.users[userID], nil
}

func (u *userRepo) GetAll(_ context.Context, limit int) ([]entity.User, error) {
	if len(u.users) < limit {
		return u.users, nil
	}
	return u.users[:limit], nil
}

func (u *userRepo) GetByEmail(_ context.Context, email string) (*entity.User, error) {
	for _, user := range u.users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, errNoSuchUser
}

func (u *userRepo) UpdateCaption(_ context.Context, userID int, caption string) error {
	for i := range u.users {
		if u.users[i].ID == userID {
			u.users[i].Caption = caption
			return nil
		}
	}
	return errNoSuchUser
}

func (u *userRepo) Delete(_ context.Context, userID int) error {
	for i := range u.users {
		if u.users[i].ID == userID {
			deletedNow := time.Now()
			u.users[i].DeletedAt = &deletedNow
			return nil
		}
	}
	return errNoSuchUser
}
