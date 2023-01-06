package entity

import "time"

type User struct {
	ID        int        `sql:"id"`
	Email     string     `sql:"email"`
	FullName  string     `sql:"full_name"`
	Caption   string     `sql:"caption"`
	BirthDate time.Time  `sql:"birth_date"`
	DeletedAt *time.Time `sql:"deleted_at"`
}
