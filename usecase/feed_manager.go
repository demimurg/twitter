package twitter

import (
	"context"
	"errors"

	"github.com/demimurg/twitter/entity"
	"github.com/demimurg/twitter/pkg/log"
)

type FeedManager interface {
	// AddFollower will subscribe follower to new tweets of the user
	AddFollower(userID, toUserID int) error
	// RemoveFollower will unsubscribe follower from new tweets of the user
	RemoveFollower(userID, fromUserID int) error

	AddNewTweet(userID int, text string) error
	GiveNewsFeed(userID int) ([]entity.Tweet, error)

	EditTweet(tweetID int, text string) error
	EditComment(commentID int, text string) error
}

var (
	ErrUserDeactivated   = errors.New("user was deactivated")
	ErrTweetLengthTooBig = errors.New("tweet text exceeded maximum length ")
)

// NewFeedManager returns usecase for work with user news feed
func NewFeedManager(
	usersRepo UserRepository, followersRepo FollowerRepository,
	tweetsRepo TweetRepository, commentsRepo CommentsRepository,
) FeedManager {
	return &feedManager{
		usersRepo: usersRepo, followersRepo: followersRepo,
		tweetsRepo: tweetsRepo, commentsRepo: commentsRepo,
	}
}

type feedManager struct {
	usersRepo     UserRepository
	followersRepo FollowerRepository
	tweetsRepo    TweetRepository
	commentsRepo  CommentsRepository
}

func (fm *feedManager) AddFollower(userID, toUserID int) error {
	return fm.followersRepo.Add(userID, toUserID)
}

func (fm *feedManager) RemoveFollower(userID, fromUserID int) error {
	return fm.followersRepo.Remove(userID, fromUserID)
}

func (fm *feedManager) AddNewTweet(userID int, text string) error {
	if len(text) > 70 {
		return ErrTweetLengthTooBig
	}

	user, err := fm.usersRepo.Get(userID)
	if err != nil {
		return err
	}
	if user.DeletedAt != nil {
		return ErrUserDeactivated
	}

	return fm.tweetsRepo.Add(userID, text)
}

func (fm *feedManager) GiveNewsFeed(userID int) ([]entity.Tweet, error) {
	following, err := fm.followersRepo.GetFollowing(userID, 10)
	if err != nil {
		return nil, err
	}

	var newsFeed []entity.Tweet
	for _, followingID := range following {
		tweets, err := fm.tweetsRepo.GetLatestFromUser(followingID, 10)
		if err != nil {
			log.Error(context.TODO(), "can't get tweets",
				"userID", userID)
			continue
		}

		newsFeed = append(newsFeed, tweets...)
	}

	return newsFeed, nil
}

func (fm *feedManager) EditTweet(tweetID int, text string) error {
	return fm.tweetsRepo.UpdateText(tweetID, text)
}

func (fm *feedManager) EditComment(commentID int, text string) error {
	return fm.commentsRepo.UpdateText(commentID, text)
}
