package seeds

import (
	"context"
	"fmt"

	"hiyoko-fiber/internal/pkg/auth/v1"
	"hiyoko-fiber/internal/pkg/ent"
	entUtil "hiyoko-fiber/internal/pkg/ent/util"
	"hiyoko-fiber/utils"
)

func UsersSeed(ctx context.Context, tx *ent.Tx) error {
	var usersInputs []ent.User
	var password, _ = auth.HashPassword("Test@0113")

	count, err := tx.User.Query().Count(ctx)
	if err != nil {
		return err
	}

	for i := count; i < count+100; i++ {
		usersInputs = append(usersInputs, ent.User{
			ID:         entUtil.NewULID(),
			OriginalID: fmt.Sprintf("user-%s+%d", utils.RandomString(5), i),
			Email:      fmt.Sprintf("user+%d@example.com", i),
			Password:   password,
		})
	}

	bulk := make([]*ent.UserCreate, len(usersInputs))
	for i, input := range usersInputs {
		bulk[i] = tx.User.
			Create().
			SetID(input.ID).
			SetOriginalID(input.OriginalID).
			SetEmail(input.Email).
			SetPassword(input.Password)
	}
	_, err = tx.User.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
