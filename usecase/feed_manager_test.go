package twitter

import (
	"strings"
	"testing"
	"time"
	"twitter/entity"
	"twitter/usecase/mock"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestFeedManager_AddNewTweet(t *testing.T) {
	t.Parallel()

	var (
		fakeUserID = "1"
	)

	type mocks struct {
		*mock.UserRepositoryMock
		*mock.TweetRepositoryMock
	}
	testCases := []struct {
		name      string
		tweetText string
		setup     func(mocks)
		wantError bool
	}{
		{
			name:      "basic case",
			tweetText: "hello world!",
			setup: func(m mocks) {
				m.UserRepositoryMock.GetMock.
					Expect(fakeUserID).
					Return(&entity.User{ID: fakeUserID}, nil)
				m.TweetRepositoryMock.AddMock.
					Expect(fakeUserID, "hello world!").
					Return(nil)
			},
		},
		{
			name:      "tweet length is to big",
			tweetText: strings.Repeat("-", 71),
			setup: func(m mocks) {
				m.UserRepositoryMock.GetMock.
					Expect(fakeUserID).
					Return(&entity.User{ID: fakeUserID}, nil)
				m.TweetRepositoryMock.AddMock.
					Expect(fakeUserID, strings.Repeat("-", 71)).
					Return(nil)
			},
		},
		{
			name:      "user was deactivated",
			tweetText: "i'm alive",
			setup: func(m mocks) {
				monthAgo := time.Now().Add(-30 * 24 * time.Hour)
				m.UserRepositoryMock.GetMock.
					Expect(fakeUserID).
					Return(&entity.User{ID: fakeUserID, DeletedAt: &monthAgo}, nil)
			},
			wantError: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mc := minimock.NewController(t)
			defer mc.Finish()
			m := mocks{
				mock.NewUserRepositoryMock(mc),
				mock.NewTweetRepositoryMock(mc),
			}
			tc.setup(m)

			var fm FeedManager // fix later
			err := fm.AddNewTweet(fakeUserID, tc.tweetText)
			assert.Equal(t, tc.wantError, err != nil, "not expected error: %v", err)
		})
	}

}
