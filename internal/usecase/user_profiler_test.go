package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/demimurg/twitter/internal/entity"
	"github.com/demimurg/twitter/internal/usecase/mock"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestUserProfiler_Register(t *testing.T) {
	var ctx = context.Background()

	testCases := []struct {
		name      string
		userName  string
		email     string
		password  string
		caption   string
		birthDate time.Time
		expect    func(upmocks)
		wantUser  *entity.User
		wantError bool
	}{
		{
			name:      "create user without problems",
			userName:  "Elon Musk",
			email:     "elon.musk@twitter.com",
			password:  "number-one",
			caption:   "billionaire, philanthropist",
			birthDate: date(1971, 06, 28),
			expect: func(m upmocks) {
				m.ScamDetectorClientMock.CheckEmailMock.
					Expect(ctx, "elon.musk@twitter.com").
					Return(nil)
				m.UserRepositoryMock.AddMock.Expect(
					ctx, "Elon Musk", "elon.musk@twitter.com", b64("number-one"),
					"billionaire, philanthropist", date(1971, 6, 28),
				).Return(1, nil)
			},
			wantUser: &entity.User{
				ID: 1, FullName: "Elon Musk", Email: "elon.musk@twitter.com",
				Caption: "billionaire, philanthropist", BirthDate: date(1971, 6, 28),
			},
		},
		{
			name:     "registration with fake email",
			userName: "Elon Musk",
			email:    "real-elon-musk@twittor.org",
			expect: func(m upmocks) {
				m.ScamDetectorClientMock.CheckEmailMock.
					Expect(ctx, "real-elon-musk@twittor.org").
					Return(ErrFakeEmail)
			},
			wantError: true,
		},
		{
			name:      "register even if scam client broken",
			userName:  "Elon Musk",
			email:     "elon.musk@twitter.com",
			password:  "number-one",
			birthDate: date(1971, 06, 28),
			expect: func(m upmocks) {
				// client may return any error, but not ErrFakeEmail
				m.ScamDetectorClientMock.CheckEmailMock.
					Expect(ctx, "elon.musk@twitter.com").
					Return(errors.New("sorry, scam client is down"))
				// problem with scam client doesn't stop us to go here
				// it's called graceful degradation
				m.UserRepositoryMock.AddMock.Expect(
					ctx, "Elon Musk", "elon.musk@twitter.com",
					b64("number-one"), "", date(1971, 6, 28),
				).Return(1, nil)
			},
			wantUser: &entity.User{ID: 1, FullName: "Elon Musk", Email: "elon.musk@twitter.com", BirthDate: date(1971, 6, 28)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			up := newUserProfiler(t, tc.expect)
			user, err := up.Register(ctx, tc.userName, tc.email, tc.password, tc.caption, tc.birthDate)
			assert.Equal(t, tc.wantError, err != nil, "not expected error: %v", err)
			assert.Equal(t, tc.wantUser, user)
		})
	}
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func TestUserProfiler_Deactivate(t *testing.T) {
	ctx := context.Background()
	expect := func(m upmocks) {
		m.UserRepositoryMock.DeleteMock.
			Expect(ctx, fakeUserID).Return(nil)
	}
	up := newUserProfiler(t, expect)

	err := up.Deactivate(ctx, fakeUserID)
	assert.NoError(t, err)
}

// ---------------------------------- SETUP MOCKS ----------------------------------

// user profiler mocks
type upmocks struct {
	*mock.UserRepositoryMock
	*mock.ScamDetectorClientMock
}

// newUserProfiler will create user profiler with mocked dependencies ans set expectations to them
func newUserProfiler(t *testing.T, expect func(upmocks)) UserProfiler {
	mc := minimock.NewController(t)
	t.Cleanup(mc.Finish)

	m := upmocks{
		mock.NewUserRepositoryMock(mc),
		mock.NewScamDetectorClientMock(mc),
	}
	expect(m)

	return NewUserProfiler(m.UserRepositoryMock, m.ScamDetectorClientMock)
}
