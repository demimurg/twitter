package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/demimurg/twitter/internal/entity"
	"github.com/demimurg/twitter/pkg/log"
)

type UserRegistrator interface {
	Register(ctx context.Context, name, email string, birthDate time.Time) (*entity.User, error)
	Deactivate(ctx context.Context, userID int) error
}

var ErrValidationFailed = errors.New("validation failed")

func NewUserRegistrator(userRepo UserRepository, scamClient ScamDetectorClient) UserRegistrator {
	return &userRegistrator{userRepo, scamClient}
}

type userRegistrator struct {
	userRepo   UserRepository
	scamClient ScamDetectorClient
}

func (ur *userRegistrator) Register(ctx context.Context, name, email string, birthDate time.Time) (*entity.User, error) {
	err := ur.scamClient.CheckEmail(ctx, email)
	if err != nil {
		if errors.Is(err, ErrFakeEmail) {
			return nil, fmt.Errorf("%s can't be registered: %w", name, err)
		}
		log.Error(ctx, "scam client returned error", err)
	}

    id, err := ur.userRepo.Add(ctx, name, email, birthDate)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID: id, Email: email,
        FullName: name, BirthDate: birthDate,
	}, nil
}

func (ur *userRegistrator) Deactivate(ctx context.Context, userID int) error {
	return ur.userRepo.Delete(ctx, userID)
}
