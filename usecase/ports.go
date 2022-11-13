package twitter

import (
	"context"
	"fmt"
	"time"

	"github.com/demimurg/twitter/entity"
)

//go:generate minimock -g -o ./mock -s .go

type UserRepository interface {
	Add(ctx context.Context, name string) error
	Get(ctx context.Context, userID int) (*entity.User, error)
	UpdateCaption(ctx context.Context, userID int, caption string) error
	Delete(ctx context.Context, userID int) error
}

type TweetRepository interface {
	Add(ctx context.Context, userID int, tweetText string) error
	UpdateText(ctx context.Context, tweetID int, newText string) error
	GetLatestFromUser(ctx context.Context, userID int, limit int) ([]entity.Tweet, error)
}

type CommentsRepository interface {
	Add(ctx context.Context, userID int, tweetID, text string) error
	UpdateText(ctx context.Context, commentID int, newText string) error
}

type FollowerRepository interface {
	// Add follower user id linked to another user
	Add(ctx context.Context, followerID, toUserID int) error
	// Remove follower from some user
	Remove(ctx context.Context, followerID, fromUserID int) error
	// GetFollowing users with topN limit
	GetFollowing(ctx context.Context, userID, topN int) ([]int, error)
	// GetFollowers give subscribed users ids
	GetFollowers(ctx context.Context, userID, topN int) ([]int, error)
}

var ErrFakeEmail = fmt.Errorf("this email is a fake")

type ScamDetectorClient interface {
	CheckEmail(ctx context.Context, email string) error
}

// ------------------------ USECASE ------------------------

type Registrator interface {
	RegisterUser(ctx context.Context, name, email string, birthDate *time.Time) (*entity.User, error)
	DeactivateUser(ctx context.Context, userID int) error
}
