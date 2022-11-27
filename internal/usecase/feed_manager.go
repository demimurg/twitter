package usecase

import (
	"context"
	"errors"
    "fmt"

    "github.com/demimurg/twitter/internal/entity"
	"github.com/demimurg/twitter/pkg/log"
)

type FeedManager interface {
	// AddFollower will subscribe follower to new tweets of the user
	AddFollower(ctx context.Context, userID, toUserID int) error
	// RemoveFollower will unsubscribe follower from new tweets of the user
	RemoveFollower(ctx context.Context, userID, fromUserID int) error

	AddNewTweet(ctx context.Context, userID int, text string) error
	GiveNewsFeed(ctx context.Context, userID int) ([]entity.Tweet, error)

	EditTweet(ctx context.Context, tweetID int, text string) error
	EditComment(ctx context.Context, commentID int, text string) error
}

var (
	ErrUserDeactivated   = errors.New("user was deactivated")
)

// NewFeedManager returns usecase for work with user news feed
func NewFeedManager(
	usersRepo UserRepository, followersRepo FollowerRepository,
	tweetsRepo TweetRepository, commentsRepo CommentsRepository,
) FeedManager {
	return &feedManager{usersRepo, followersRepo, tweetsRepo, commentsRepo}
}

type feedManager struct {
	usersRepo     UserRepository
	followersRepo FollowerRepository
	tweetsRepo    TweetRepository
	commentsRepo  CommentsRepository
}

func (fm *feedManager) AddFollower(ctx context.Context, userID, toUserID int) error {
	return fm.followersRepo.Add(ctx, userID, toUserID)
}

func (fm *feedManager) RemoveFollower(ctx context.Context, userID, fromUserID int) error {
	return fm.followersRepo.Remove(ctx, userID, fromUserID)
}

func (fm *feedManager) AddNewTweet(ctx context.Context, userID int, text string) error {
	if len(text) > 70 {
        return fmt.Errorf(
            "%w: tweet length %d more than allowed %d",
            ErrValidationFailed, len(text), 70,
        )
	}

	user, err := fm.usersRepo.Get(ctx, userID)
	if err != nil {
		return err
	}
	if user.DeletedAt != nil {
		return ErrUserDeactivated
	}

	return fm.tweetsRepo.Add(ctx, userID, text)
}

func (fm *feedManager) GiveNewsFeed(ctx context.Context, userID int) ([]entity.Tweet, error) {
	following, err := fm.followersRepo.GetFollowing(ctx, userID, 10)
	if err != nil {
		return nil, err
	}

	var newsFeed []entity.Tweet
	for _, followingID := range following {
		tweets, err := fm.tweetsRepo.GetLatestFromUser(ctx, followingID, 10)
		if err != nil {
			log.Error(ctx, "can't get tweets",
				"error", err,
				"userID", userID)
			continue
		}

		newsFeed = append(newsFeed, tweets...)
	}

	return newsFeed, nil
}

func (fm *feedManager) EditTweet(ctx context.Context, tweetID int, text string) error {
	return fm.tweetsRepo.UpdateText(ctx, tweetID, text)
}

func (fm *feedManager) EditComment(ctx context.Context, commentID int, text string) error {
	return fm.commentsRepo.UpdateText(ctx, commentID, text)
}
