package twitter

import "twitter/entity"

type FeedManager interface {
	AddNewTweet(userID string, text string) error
	GiveNewsFeed(userID string) (*entity.NewsFeed, error)
	EditTweet(messageID, text string) error
	EditComment(commentID, text string) error
}

type feedManager struct {
}

func (f *feedManager) AddNewTweet(userID string, text string) error {
	return nil
}

func (f *feedManager) GiveNewsFeed(userID string) (*entity.NewsFeed, error) {
	return nil, nil
}

func (f *feedManager) EditTweet(messageID, text string) error {
	return nil
}

func (f *feedManager) EditComment(commentID, text string) error {
	return nil
}
