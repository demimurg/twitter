package entity

import "time"

// MaxAllowedSymbols set max length for all tweets, be short and smart
const MaxAllowedSymbols = 140

type Tweet struct {
	ID         int    `sql:"id"`
	UserID     int    `sql:"user_id"`
	Text       string `sql:"text"`
	Likes      int    `sql:"likes"`
	CommentIDs []string
	CreatedAt  time.Time `sql:"created_at"`
}

type Comment struct {
	ID     int    `sql:"id"`
	UserID int    `sql:"user_id"`
	Text   string `sql:"text"`
	Likes  int    `sql:"likes"`
}
