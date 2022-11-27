package entity

import "time"

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
