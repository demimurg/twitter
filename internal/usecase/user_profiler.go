package usecase

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/demimurg/twitter/internal/entity"
	"github.com/demimurg/twitter/pkg/log"
)

type UserProfiler interface {
	Register(ctx context.Context, name, email, password, caption string, birthDate time.Time) (*entity.User, error)
	Deactivate(ctx context.Context, userID int) error
	UpdateCaption(ctx context.Context, userID int, newCaption string) error

	Login(ctx context.Context, email, password string) (*entity.User, error)
}

func NewUserProfiler(userRepo UserRepository, scamClient ScamDetectorClient) UserProfiler {
	return &userProfiler{userRepo, scamClient}
}

type userProfiler struct {
	userRepo   UserRepository
	scamClient ScamDetectorClient
}

func (up *userProfiler) Register(ctx context.Context, name, email, password, caption string, birthDate time.Time) (*entity.User, error) {
	err := up.scamClient.CheckEmail(ctx, email)
	if err != nil {
		if errors.Is(err, ErrFakeEmail) {
			return nil, fmt.Errorf("%s can't be registered: %w", name, err)
		}
		log.Error(ctx, "scam client returned error", err)
	}

	id, err := up.userRepo.Add(ctx, name, email, b64(password), caption, birthDate)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID: id, Email: email, Caption: caption,
		FullName: name, BirthDate: birthDate,
	}, nil
}

func (up *userProfiler) Login(ctx context.Context, email, password string) (*entity.User, error) {
	return up.userRepo.GetByEmail(ctx, email, b64(password))
}

func (up *userProfiler) Deactivate(ctx context.Context, userID int) error {
	return up.userRepo.Delete(ctx, userID)
}

func (up *userProfiler) UpdateCaption(ctx context.Context, userID int, newCaption string) error {
	return up.userRepo.UpdateCaption(ctx, userID, newCaption)
}

// b64 shorthand for standart base64 encoding
func b64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
