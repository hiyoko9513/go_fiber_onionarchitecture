package repository

import (
	"context"

	"hiyoko-fiber/internal/domain/services"
	"hiyoko-fiber/internal/infrastructure/database"
	"hiyoko-fiber/internal/pkg/ent"
)

type userRepository struct {
	conn *database.EntClient
}

func NewUserRepository(conn *database.EntClient) services.UserRepository {
	return &userRepository{conn}
}

func (r *userRepository) Get(ctx context.Context, id string) (*ent.User, error) {
	u, err := r.conn.User.Get(ctx, id)
	return u, err
}

func (r *userRepository) Create(ctx context.Context, u *ent.User) (*ent.User, error) {
	u, err := r.conn.User.Create().
		SetID(u.ID).
		SetEmail(u.Email).
		SetPassword(u.Password).
		Save(ctx)
	return u, err
}

func (r *userRepository) Update(ctx context.Context, u *ent.User) (*ent.User, error) {
	u, err := u.Update().SetEmail(u.Email).Save(ctx)
	return u, err
}
