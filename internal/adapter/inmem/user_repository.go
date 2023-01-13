package inmem

import (
	"context"
    "github.com/pkg/errors"
    "time"

	"github.com/demimurg/twitter/internal/entity"
	"github.com/demimurg/twitter/internal/usecase"
)

var errNoSuchUser = errors.New("there is no such user in memory")

func NewUserRepository() usecase.UserRepository {
    s := make([]entity.User, 0, 100) // initial storage size 100 users
    s = append(s, entity.User{FullName: "Pavel Durov"}) // book place to start new ids from 1
	return &userRepo{storage: s}
}

type userRepo struct {
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
        return nil, errNoSuchUser
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

    return nil, errNoSuchUser
}

func (u *userRepo) UpdateCaption(_ context.Context, userID int, caption string) error {
    for i := range u.storage {
        if u.storage[i].ID == userID {
            u.storage[i].Caption = caption
            return nil
        }
    }
    return errNoSuchUser
}

func (u *userRepo) Delete(_ context.Context, userID int) error {
    for i := range u.storage {
        if u.storage[i].ID == userID {
            deletedNow := time.Now()
            u.storage[i].DeletedAt = &deletedNow
            return nil
        }
    }
    return errNoSuchUser
}