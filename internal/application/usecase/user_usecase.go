package usecase

import (
	"context"

	"hiyoko-fiber/internal/domain/entities/users"
	"hiyoko-fiber/internal/domain/services"
	"hiyoko-fiber/internal/pkg/ent"
	"hiyoko-fiber/internal/presentation/http/app/input"
	"hiyoko-fiber/pkg/logging/file"
)

type UserUseCase interface {
	//GetUser(ctx context.Context, id string) (*ent.User, error)
	Signup(ctx context.Context, input *input.UserCreateInput) (*users.UserEntity, error)
	//UpdateUser(ctx context.Context, id string) (*ent.User, error)
}

type userUseCase struct {
	services.UserRepository
}

func NewUserUseCase(r services.UserRepository) UserUseCase {
	return &userUseCase{r}
}

//func (u *userUseCase) GetUser(ctx context.Context, id string) (*ent.User, error) {
//	return u.UserRepository.Get(ctx, id)
//}

func (u *userUseCase) Signup(ctx context.Context, input *input.UserCreateInput) (*users.UserEntity, error) {
	user := &users.UserEntity{
		ID:       input.ID,
		Email:    input.Email,
		Password: input.Password,
	}

	entUser, err := u.UserRepository.Create(ctx, &ent.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		logger.Error("Error creating user", "input", user, "error", err)
		return &users.UserEntity{}, err
	}

	user.Sub = string(entUser.Sub)
	user.CreatedAt = entUser.CreatedAt
	user.UpdatedAt = entUser.UpdatedAt

	return user, nil
}

//func (u *userUseCase) UpdateUser(ctx context.Context, id string) (*ent.User, error) {
//	user, err := u.UserRepository.Get(ctx, id)
//	if err != nil {
//		return nil, err
//	}
//	return u.UserRepository.Update(ctx, user)
//}
