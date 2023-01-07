package entity

import "time"

// MaxAllowedSymbols set max length for all tweets, be short and smart
const MaxAllowedSymbols = 140

type Tweet struct {
	ID        int
	UserID    int
	Text      string
	Likes     int
	CreatedAt time.Time
}

type Comment struct {
	ID     int
	UserID int
	Text   string
	Likes  int
}
