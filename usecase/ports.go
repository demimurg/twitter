package twitter

import (
	"fmt"
	"time"

	"github.com/demimurg/twitter/entity"
)

//go:generate minimock -g -o ./mock -s .go

type UserRepository interface {
	Add(name string) error
	Get(userID int) (*entity.User, error)
	UpdateCaption(userID int, caption string) error
	Delete(userID int) error
}

type TweetRepository interface {
	Add(userID int, tweetText string) error
	UpdateText(tweetID int, newText string) error
	GetLatestFromUser(userID int, limit int) ([]entity.Tweet, error)
}

type CommentsRepository interface {
	Add(userID int, tweetID, text string) error
	UpdateText(commentID int, newText string) error
}

type FollowerRepository interface {
	// Add follower user id linked to another user
	Add(followerID, toUserID int) error
	// Remove follower from some user
	Remove(followerID, fromUserID int) error
	// GetFollowing users with topN limit
	GetFollowing(userID, topN int) ([]int, error)
	// GetFollowers give subscribed users ids
	GetFollowers(userID, topN int) ([]int, error)
}

var ErrFakeEmail = fmt.Errorf("this email is a fake")

type ScamDetectorClient interface {
	CheckEmail(email string) error
}

// ------------------------ USECASE ------------------------

type Registrator interface {
	RegisterUser(name, email string, birthDate *time.Time) (*entity.User, error)
	DeactivateUser(userID int) error
}
