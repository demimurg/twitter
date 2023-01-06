package inmem

import (
	"context"
	"fmt"
	"time"

	"github.com/demimurg/twitter/internal/entity"
	"github.com/demimurg/twitter/internal/usecase"
)

func NewUserRepository() usecase.UserRepository {
	return &userRepo{storage: make([]entity.User, 0, 100)} // initial storage size 100
}

type userRepo struct {
	usecase.UserRepository
	storage []entity.User
}

func (u *userRepo) Add(_ context.Context, name, email, caption string, birthDate time.Time) (int, error) {
	id := len(u.storage)
	u.storage = append(u.storage, entity.User{
		ID:        id,
		Email:     email,
		FullName:  name,
		Caption:   caption,
		BirthDate: birthDate,
	})

	return id, nil
}

func (u *userRepo) Get(_ context.Context, userID int) (*entity.User, error) {
	if userID >= len(u.storage) || userID < 0 {
		return nil, fmt.Errorf("there is no such user id %d", userID)
	}

	return &u.storage[userID], nil
}

func (u *userRepo) GetAll(_ context.Context, limit int) ([]entity.User, error) {
	if len(u.storage) < limit {
		return u.storage, nil
	}
	return u.storage[:limit], nil
}

func (u *userRepo) GetByEmail(_ context.Context, email string) (*entity.User, error) {
	for _, user := range u.storage {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("there is no such user email %q", email)
}
