package usecase

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/demimurg/twitter/internal/entity"
	"github.com/demimurg/twitter/internal/usecase/mock"

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
		expect    func(fmmocks)
		wantError bool
	}{
		{
			name:      "basic case",
			tweetText: "hello world!",
			expect: func(m fmmocks) {
				m.UserRepositoryMock.GetMock.
					Expect(ctx, fakeUserID).
					Return(&entity.User{ID: fakeUserID}, nil)
				m.TweetRepositoryMock.AddMock.
					Expect(ctx, fakeUserID, "hello world!").
					Return(0, nil)
			},
		},
		{
			name:      "user was deactivated",
			tweetText: "i'm alive",
			expect: func(m fmmocks) {
				monthAgo := time.Now().Add(-30 * 24 * time.Hour)
				m.UserRepositoryMock.GetMock.
					Expect(ctx, fakeUserID).
					Return(&entity.User{ID: fakeUserID, DeletedAt: &monthAgo}, nil)
			},
			wantError: true,
		},
		{
			name:      "tweet length is too big",
			tweetText: strings.Repeat("-", entity.MaxAllowedSymbols+1),
			expect: func(m fmmocks) {
				// there is no calls to upmocks, will be entry validation of length and error returns
			},
			wantError: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			feedManager := newFeedManager(t, tc.expect)
			_, err := feedManager.AddTweet(ctx, fakeUserID, tc.tweetText)

			assert.Equal(t, tc.wantError, err != nil, "not expected error: %v", err)
		})
	}
}

func TestFeedManager_GiveNewsFeed(t *testing.T) {
	testCases := []struct {
		name      string
		expect    func(fmmocks)
		wantFeed  []entity.Tweet
		wantError bool
	}{
		{
			name: "just give my newsfeed",
			expect: func(m fmmocks) {
				// get following users
				friend1, friend2 := 1, 2
				m.FollowerRepositoryMock.GetFolloweeMock.
					Expect(ctx, fakeUserID, 10).
					Return([]int{friend1, friend2}, nil)
				// fetch their latest tweets
				m.TweetRepositoryMock.
					GetLatestMock.
					When(ctx, friend1, 10).
					Then([]entity.Tweet{{Text: "wake up"}, {Text: "eat"}}, nil).
					GetLatestMock.
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
			expect: func(m fmmocks) {
				m.FollowerRepositoryMock.GetFolloweeMock.
					Expect(ctx, fakeUserID, 10).
					Return(nil, errors.New("sorry guys, i'have drop off database"))
			},
			wantError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			feedManager := newFeedManager(t, tc.expect)
			feed, err := feedManager.GetNewsFeed(ctx, fakeUserID)

			assert.Equal(t, tc.wantError, err != nil, "not expected error: %v", err)
			assert.Equal(t, tc.wantFeed, feed)
		})
	}
}

// ---------------------------------- SETUP MOCKS ----------------------------------

// feed manager mocks
type fmmocks struct {
	*mock.UserRepositoryMock
	*mock.FollowerRepositoryMock
	*mock.TweetRepositoryMock
}

// newFeedManager will create feed manager with mocked dependencies ans set expectations to them
func newFeedManager(t *testing.T, expect func(fmmocks)) FeedManager {
	mc := minimock.NewController(t)
	t.Cleanup(mc.Finish)

	m := fmmocks{
		mock.NewUserRepositoryMock(mc),
		mock.NewFollowerRepositoryMock(mc),
		mock.NewTweetRepositoryMock(mc),
	}
	expect(m)

	return NewFeedManager(m.UserRepositoryMock, m.FollowerRepositoryMock, m.TweetRepositoryMock)
}
