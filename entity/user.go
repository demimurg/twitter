package entity

import "time"

type User struct {
	ID            string
	Login         string
	Name          string
	Caption       string
	DeletedAt     *time.Time
	SubscriberIDs []string
}
