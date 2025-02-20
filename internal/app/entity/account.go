package entity

import "time"

type Account struct {
	ID        uint32
	CreatedAt time.Time
	UpdatedAt time.Time
	UUID      string
	Token     string
}
