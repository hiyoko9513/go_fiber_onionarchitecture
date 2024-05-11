package users

import "time"

type UserEntity struct {
	ID         string
	Status     Status
	OriginalID string
	Email      string
	Password   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Status int8

const (
	StatusInactive   Status = 0 // inactive
	StatusActive     Status = 1 // active
	StatusUnverified Status = 2 // unverified
)

func (s Status) Default() Status {
	return StatusUnverified
}

func (s Status) ToInt8() int8 {
	return int8(s)
}
