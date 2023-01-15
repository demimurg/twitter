package inmem

import (
	"context"
	"errors"
	"sort"
	"time"

	"github.com/demimurg/twitter/internal/entity"
	"github.com/demimurg/twitter/internal/usecase"
)

var (
	errTweetNotFound   = errors.New("can't find tweet in memory")
	errCommentNotFound = errors.New("can't find comment in memory")
)

func NewTweetRepository() usecase.TweetRepository {
	return &tweetRepo{
		tweets:   make([]entity.Tweet, 0, 1000),
		comments: make([]comment, 0, 1000),
	}
}

type tweetRepo struct {
	tweets   []entity.Tweet
	comments []comment
}

type comment struct {
	entity.Comment
	tweetID int
}

func (t *tweetRepo) Add(_ context.Context, userID int, tweetText string) (id int, err error) {
	id = len(t.tweets) // add to the end
	t.tweets = append(t.tweets, entity.Tweet{
		ID:        id,
		UserID:    userID,
		Text:      tweetText,
		CreatedAt: time.Now().UTC(),
	})
	return id, nil
}

func (t *tweetRepo) Update(_ context.Context, tweetID int, newText string) error {
	for i := range t.tweets {
		if t.tweets[i].ID == tweetID {
			t.tweets[i].Text = newText
			return nil
		}
	}
	return errTweetNotFound
}

func (t *tweetRepo) AddComment(_ context.Context, userID, tweetID int, commentText string) (id int, err error) {
	var c comment
	c.ID = len(t.comments)
	c.UserID = userID
	c.tweetID = tweetID
	c.Text = commentText

	t.comments = append(t.comments, c)
	return c.ID, nil
}

func (t *tweetRepo) UpdateComment(_ context.Context, commentID int, newText string) error {
	for i := range t.comments {
		if t.comments[i].ID == commentID {
			t.comments[i].Text = newText
			return nil
		}
	}
	return errCommentNotFound
}

func (t *tweetRepo) GetLatest(_ context.Context, userID int, limit int) ([]entity.Tweet, error) {
	var userTweets []entity.Tweet
	for _, tweet := range t.tweets {
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
