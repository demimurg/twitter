package usecase

import (
	"context"
	"errors"
	"github.com/demimurg/twitter/internal/entity"
	"github.com/demimurg/twitter/internal/usecase/mock"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserRegistrator_Register(t *testing.T) {
	var ctx = context.Background()

	testCases := []struct {
		name      string
		userName  string
		email     string
		birthDate string
		expect    func(urmocks)
		wantUser  *entity.User
		wantError bool
	}{
		{
			name:      "create user without problems",
			userName:  "Elon Musk",
			email:     "elon.musk@twitter.com",
			birthDate: "1971-06-28",
			expect: func(m urmocks) {
				m.ScamDetectorClientMock.CheckEmailMock.
					Expect(ctx, "elon.musk@twitter.com").
					Return(nil)
				m.UserRepositoryMock.AddMock.
					Expect(ctx, "Elon Musk", "elon.musk@twitter.com", date(1971, 6, 28)).
					Return(1, nil)
			},
			wantUser: &entity.User{ID: 1, FullName: "Elon Musk", Email: "elon.musk@twitter.com", BirthDate: date(1971, 6, 28)},
		},
		{
			name:      "birth date is not valid",
			userName:  "Elon Musk",
			email:     "elon.musk@twitter.com",
			birthDate: "28.06.1971 birth", // wrong format
			expect:    func(urmocks) {},
			wantError: true,
		},
		{
			name:      "registration with fake email",
			userName:  "Elon Musk",
			email:     "real-elon-musk@twittor.org",
			birthDate: "1971-06-28",
			expect: func(m urmocks) {
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
			birthDate: "1971-06-28",
			expect: func(m urmocks) {
				// client may return any error, but not ErrFakeEmail
				m.ScamDetectorClientMock.CheckEmailMock.
					Expect(ctx, "elon.musk@twitter.com").
					Return(errors.New("sorry, scam client is down"))
				// problem with scam client doesn't stop us to go here
                // it's called gracefull degradation
				m.UserRepositoryMock.AddMock.
					Expect(ctx, "Elon Musk", "elon.musk@twitter.com", date(1971, 6, 28)).
					Return(1, nil)
			},
			wantUser: &entity.User{ID: 1, FullName: "Elon Musk", Email: "elon.musk@twitter.com", BirthDate: date(1971, 6, 28)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ur := newUserRegistrator(t, tc.expect)
			user, err := ur.Register(ctx, tc.userName, tc.email, tc.birthDate)
			assert.Equal(t, tc.wantError, err != nil, "not expected error: %v", err)
			assert.Equal(t, tc.wantUser, user)
		})
	}
}

func date(year, month, day int) *time.Time {
	d := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return &d
}

func TestUserRegistrator_Deactivate(t *testing.T) {
	ctx := context.Background()
	expect := func(m urmocks) {
		m.UserRepositoryMock.DeleteMock.
			Expect(ctx, fakeUserID).Return(nil)
	}
	r := newUserRegistrator(t, expect)

	err := r.Deactivate(ctx, fakeUserID)
    assert.NoError(t, err)
}

// ---------------------------------- SETUP MOCKS ----------------------------------

// user registrator fmmocks
type urmocks struct {
	*mock.UserRepositoryMock
	*mock.ScamDetectorClientMock
}

// newUserRegistrator will create user registrator with mocked dependencies ans set expectations to them
func newUserRegistrator(t *testing.T, expect func(urmocks)) UserRegistrator {
	mc := minimock.NewController(t)
	t.Cleanup(mc.Finish)

	m := urmocks{
		mock.NewUserRepositoryMock(mc),
		mock.NewScamDetectorClientMock(mc),
	}
	expect(m)

	return NewUserRegistrator(m.UserRepositoryMock, m.ScamDetectorClientMock)
}
