package inmem

import (
	"context"
	"github.com/demimurg/twitter/internal/entity"
	"github.com/demimurg/twitter/internal/usecase"
	"sort"
	"time"
)

func NewTweetRepository() usecase.TweetRepository {
	return &tweetRepo{storage: make([]entity.Tweet, 0, 1000)} // initial storage size
}

type tweetRepo struct {
	usecase.TweetRepository
	storage []entity.Tweet
}

func (t *tweetRepo) Add(_ context.Context, userID int, tweetText string) error {
	t.storage = append(t.storage, entity.Tweet{
		ID:        len(t.storage), // add to the end
		UserID:    userID,
		Text:      tweetText,
		CreatedAt: time.Now().UTC(),
	})
	return nil
}

func (t *tweetRepo) GetLatestFromUser(_ context.Context, userID int, limit int) ([]entity.Tweet, error) {
	var userTweets []entity.Tweet
	for _, tweet := range t.storage {
		if tweet.UserID == userID {
			userTweets = append(userTweets, tweet)
		}
	}

	// recently tweets first
	sort.Slice(userTweets, func(i, j int) bool {
		return userTweets[i].CreatedAt.After(userTweets[j].CreatedAt)
	})

    if limit > len(userTweets) {
       return userTweets, nil
	}
    return userTweets[:limit], nil
}
