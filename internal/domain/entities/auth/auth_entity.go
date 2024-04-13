package auth

import (
	"hiyoko-fiber/internal/domain/entities/users"
)

type AuthenticationEntity struct {
	Token string
	Exp   int64
	User  users.UserEntity
}
