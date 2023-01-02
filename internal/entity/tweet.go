package entity

import "time"

// MaxAllowedSymbols set max length for all tweets, be short and smart
const MaxAllowedSymbols = 140

type Tweet struct {
	ID, UserID int
	Text       string
	Likes      int
	CommentIDs []string
	CreatedAt  time.Time
}

type Comment struct {
	ID, UserID int
	Text       string
	Likes      int
}
