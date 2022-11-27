package inmem

import (
	"context"
    "fmt"
    "github.com/demimurg/twitter/internal/entity"
	"github.com/demimurg/twitter/internal/usecase"
	"time"
)

func NewUserRepository() usecase.UserRepository {
	return &userRepo{storage: make([]*entity.User, 0, 100)} // initial storage size 100
}

type userRepo struct {
	usecase.UserRepository
	storage []*entity.User
}

func (u *userRepo) Add(_ context.Context, name, email string, birthDate *time.Time) (int, error) {
	id := len(u.storage)
	u.storage = append(u.storage, &entity.User{
		ID:        id,
		Email:     email,
		FullName:  name,
		BirthDate: birthDate,
	})

	return id, nil
}

func (u *userRepo) Get(_ context.Context, userID int) (*entity.User, error) {
    if userID >= len(u.storage) {
        return nil, fmt.Errorf("there is no such user id %d", userID)
    }
    
    return u.storage[userID], nil
}
