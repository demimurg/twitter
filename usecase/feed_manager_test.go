package twitter

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/demimurg/twitter/entity"
	"github.com/demimurg/twitter/usecase/mock"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

var (
	fakeUserID = 666
	ctx        = context.Background()
)

func TestFeedManager_AddNewTweet(t *testing.T) {
	testCases := []struct {
		name      string
		tweetText string
		expect    func(mocks)
		wantError bool
	}{
		{
			name:      "basic case",
			tweetText: "hello world!",
			expect: func(m mocks) {
				m.UserRepositoryMock.GetMock.
					Expect(ctx, fakeUserID).
					Return(&entity.User{ID: fakeUserID}, nil)
				m.TweetRepositoryMock.AddMock.
					Expect(ctx, fakeUserID, "hello world!").
					Return(nil)
			},
		},
		{
			name:      "user was deactivated",
			tweetText: "i'm alive",
			expect: func(m mocks) {
				monthAgo := time.Now().Add(-30 * 24 * time.Hour)
				m.UserRepositoryMock.GetMock.
					Expect(ctx, fakeUserID).
					Return(&entity.User{ID: fakeUserID, DeletedAt: &monthAgo}, nil)
			},
			wantError: true,
		},
		{
			name:      "tweet length is to big",
			tweetText: strings.Repeat("-", 71),
			expect: func(m mocks) {
				// there is no calls to mocks, will be entry validation of length and error returns
			},
			wantError: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			feedManager := newFeedManager(t, tc.expect)
			err := feedManager.AddNewTweet(ctx, fakeUserID, tc.tweetText)

			assert.Equal(t, tc.wantError, err != nil, "not expected error: %v", err)
		})
	}
}

func TestFeedManager_GiveNewsFeed(t *testing.T) {
	testCases := []struct {
		name      string
		expect    func(mocks)
		wantFeed  []entity.Tweet
		wantError bool
	}{
		{
			name: "just give my newsfeed",
			expect: func(m mocks) {
				// get following users
				friend1, friend2 := 1, 2
				m.FollowerRepositoryMock.GetFollowingMock.
					Expect(ctx, fakeUserID, 10).
					Return([]int{friend1, friend2}, nil)
				// fetch their latest tweets
				m.TweetRepositoryMock.
					GetLatestFromUserMock.
					When(ctx, friend1, 10).
					Then([]entity.Tweet{{Text: "wake up"}, {Text: "eat"}}, nil).
					GetLatestFromUserMock.
					When(ctx, friend2, 10).
					Then([]entity.Tweet{{Text: "yoga class"}, {Text: "kitten"}}, nil)
				// merge them
			},
			wantFeed: []entity.Tweet{
				{Text: "wake up"}, {Text: "eat"},
				{Text: "yoga class"}, {Text: "kitten"},
			},
		},
		{
			name: "database with followers is not available",
			expect: func(m mocks) {
				m.FollowerRepositoryMock.GetFollowingMock.
					Expect(ctx, fakeUserID, 10).
					Return(nil, errors.New("sorry guys, i'have drop off database"))
			},
			wantError: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			feedManager := newFeedManager(t, tc.expect)
			feed, err := feedManager.GiveNewsFeed(ctx, fakeUserID)

			assert.Equal(t, tc.wantError, err != nil, "not expected error: %v", err)
			assert.Equal(t, tc.wantFeed, feed)
		})
	}
}

// ---------------------------------- SETUP MOCKS ----------------------------------

type mocks struct {
	*mock.UserRepositoryMock
	*mock.FollowerRepositoryMock
	*mock.TweetRepositoryMock
	*mock.CommentsRepositoryMock
}

// newFeedManager will create feed manager with mocked dependencies ans set expectations to them
func newFeedManager(t *testing.T, expect func(mocks)) FeedManager {
	mc := minimock.NewController(t)
	t.Cleanup(mc.Finish)

	m := mocks{
		mock.NewUserRepositoryMock(mc),
		mock.NewFollowerRepositoryMock(mc),
		mock.NewTweetRepositoryMock(mc),
		mock.NewCommentsRepositoryMock(mc),
	}
	expect(m)

	return NewFeedManager(
		m.UserRepositoryMock, m.FollowerRepositoryMock,
		m.TweetRepositoryMock, m.CommentsRepositoryMock,
	)
}
