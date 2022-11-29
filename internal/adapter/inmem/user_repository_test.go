package inmem

import (
	"context"
	"github.com/demimurg/twitter/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserRepository(t *testing.T) {
	ctx := context.Background()
	repo := NewUserRepository()

	elonDR := time.Date(1971, time.June, 28, 0, 0, 0, 0, time.UTC)
	fakeElon := entity.User{FullName: "Elon Musk", Email: "elon@musk.com", BirthDate: &elonDR}

	t.Run("can add and get user", func(t *testing.T) {
		wantUser := fakeElon // copy for subtest

		userID, err := repo.Add(ctx, fakeElon.FullName, fakeElon.Email, wantUser.BirthDate)
		require.NoError(t, err)
		wantUser.ID = userID

		user, err := repo.Get(ctx, wantUser.ID)
		require.NoError(t, err)
		assert.Equal(t, wantUser, *user)
	})

	t.Run("get user with wrong id", func(t *testing.T) {
		_, err := repo.Get(ctx, 100)
		require.Error(t, err)
	})
}
