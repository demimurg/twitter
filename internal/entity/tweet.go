package entity

import (
	"errors"
	"fmt"
	"time"
)

// MaxAllowedSymbols set max length for all tweets, be short and smart
const MaxAllowedSymbols = 140

var ErrValidationFailed = errors.New("validation failed")

func ValidateTweet(text string) error {
	if len(text) > MaxAllowedSymbols {
		return fmt.Errorf(
			"tweet size bigger than allowed (%d symbols): %w",
			MaxAllowedSymbols, ErrValidationFailed,
		)
	}
	return nil
}

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
