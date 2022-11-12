package twitter

import (
	"fmt"
	"time"

	"github.com/demimurg/twitter/entity"
)

//go:generate minimock -g -o ./mock -s .go

type UserRepository interface {
	Add(name string) error
	Get(userID string) (*entity.User, error)
	UpdateCaption(userID string, caption string) error
	Delete(userID string) error
}

type TweetRepository interface {
	Add(userID, tweetText string) error
}

var ErrFakeEmail = fmt.Errorf("this email is a fake")

type ScamDetectorClient interface {
	CheckEmail(email string) error
}

// ------------------------ USECASE ------------------------

type Registrator interface {
	RegisterUser(name, email string, birthDate *time.Time) (*entity.User, error)
	DeactivateUser(userID string) error
}
