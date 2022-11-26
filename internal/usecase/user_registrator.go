package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/demimurg/twitter/internal/entity"
	"github.com/demimurg/twitter/pkg/log"
	"time"
)

type UserRegistrator interface {
	Register(ctx context.Context, name, email, birthDate string) (*entity.User, error)
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

func (ur *userRegistrator) Register(ctx context.Context, name, email, birthDate string) (*entity.User, error) {
	birth, err := time.Parse("2006-01-02", birthDate) // "year-month-day" format
	if err != nil {
		return nil, fmt.Errorf("%w: expect birth date in format '2006-01-02' (year-month-day)", ErrValidationFailed)
	}

	err = ur.scamClient.CheckEmail(ctx, email)
	if err != nil {
		if errors.Is(err, ErrFakeEmail) {
			return nil, fmt.Errorf("%s can't be registered: %w", name, err)
		}
		log.Error(ctx, "scam client returned error", err)
	}

	id, err := ur.userRepo.Add(ctx, name, email, &birth)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID: id, Email: email,
		FullName: name, BirthDate: &birth,
	}, nil
}

func (ur *userRegistrator) Deactivate(ctx context.Context, userID int) error {
	return ur.userRepo.Delete(ctx, userID)
}
