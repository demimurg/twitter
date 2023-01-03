package entity

import "time"

type User struct {
	ID        int
	Email     string
	FullName  string
	Caption   string
	Followers int
	BirthDate time.Time
	DeletedAt *time.Time
}
