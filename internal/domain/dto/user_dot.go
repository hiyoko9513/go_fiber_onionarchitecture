package dto

import "time"

// todo createAtを固有のtypeに変更する（util.time.goに作成するはず？）
type UserCreationDot struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
