package inmem

import (
	"context"
	"github.com/demimurg/twitter/internal/usecase"
)

func NewFollowerRepository() usecase.FollowerRepository {
	return &followerRepo{
		followersStorage: make(map[int][]int, 100),
		followingStorage: make(map[int][]int, 100),
	}
}

type followerRepo struct {
	usecase.FollowerRepository
	followersStorage map[int][]int
	followingStorage map[int][]int
}

func (f *followerRepo) Add(_ context.Context, followerID, toUserID int) error {
	f.followersStorage[toUserID] = append(f.followersStorage[toUserID], followerID)
	f.followingStorage[followerID] = append(f.followingStorage[followerID], toUserID)
	return nil
}

func (f *followerRepo) GetFollowing(_ context.Context, userID, topN int) ([]int, error) {
	following, ok := f.followingStorage[userID]
	if !ok {
		return []int{}, nil
	}

	if len(following) < topN {
		return following, nil
	}
	return following[:topN], nil
}

func (f *followerRepo) GetFollowers(_ context.Context, userID, topN int) ([]int, error) {
	followers, ok := f.followersStorage[userID]
	if !ok {
		return []int{}, nil
	}

	if len(followers) < topN {
		return followers, nil
	}
	return followers[:topN], nil
}
