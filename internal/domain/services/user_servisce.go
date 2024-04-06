package services

import (
	"context"

	"hiyoko-fiber/internal/pkg/ent"
)

type UserRepository interface {
	Get(ctx context.Context, id string) (*ent.User, error)
	Create(ctx context.Context, user *ent.User) (*ent.User, error)
	Update(ctx context.Context, user *ent.User) (*ent.User, error)
}
