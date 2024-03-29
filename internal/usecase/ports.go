package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/demimurg/twitter/internal/entity"
)

var (
	ErrUserExists       = errors.New("user already exists")
	ErrWrongCredentials = errors.New("have no user with specified email and password")
)

type UserRepository interface {
	// Add will creates new user in repo and returns id assigned to it
	Add(ctx context.Context, name, email, password, caption string, birthDate time.Time) (id int, err error)
	Get(ctx context.Context, userID int) (*entity.User, error)
	GetAll(ctx context.Context, limit int) ([]entity.User, error)
	// GetByEmail will return ErrWrongCredentials if no user found
	GetByEmail(ctx context.Context, email, password string) (*entity.User, error)
	UpdateCaption(ctx context.Context, userID int, caption string) error
	Delete(ctx context.Context, userID int) error
}

type TweetRepository interface {
	Add(ctx context.Context, userID int, tweetText string) (id int, err error)
	Update(ctx context.Context, tweetID int, newText string) error
	AddComment(ctx context.Context, userID, tweetID int, commentText string) (id int, err error)
	UpdateComment(ctx context.Context, commentID int, newText string) error
	GetLatest(ctx context.Context, userID int, limit int) ([]entity.Tweet, error)
}

type FollowerRepository interface {
	// Add follower user id linked to another user
	Add(ctx context.Context, followerID, toUserID int) error
	// Remove follower from some user
	Remove(ctx context.Context, followerID, fromUserID int) error
	// 	GetFollowing users with topN limit
	GetFollowing(ctx context.Context, userID, topN int) ([]int, error)
	// GetFollowers give subscribed users ids
	GetFollowers(ctx context.Context, userID, topN int) ([]int, error)
}

var ErrFakeEmail = errors.New("this email is a fake")

type ScamDetectorClient interface {
	CheckEmail(ctx context.Context, email string) error
}
