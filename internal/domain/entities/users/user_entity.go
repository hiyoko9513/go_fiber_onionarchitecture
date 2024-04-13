package users

import "time"

type UserEntity struct {
	ID         string
	OriginalID string
	Email      string
	Password   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
