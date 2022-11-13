package entity

import "time"

type User struct {
	ID        int
	Login     string
	FullName  string
	Caption   string
	Followers int
	DeletedAt *time.Time
}
