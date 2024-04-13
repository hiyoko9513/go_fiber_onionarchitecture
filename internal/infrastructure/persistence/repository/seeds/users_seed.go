package seeds

import (
	"context"
	"fmt"

	"hiyoko-fiber/internal/pkg/ent"
	entUtil "hiyoko-fiber/internal/pkg/ent/util"
)

func UsersSeed(ctx context.Context, tx *ent.Tx) error {
	var usersInputs []ent.User
	for i := 0; i < 100; i++ {
		usersInputs = append(usersInputs, ent.User{
			ID: entUtil.NewULID(),
			//OriginalID: fmt.Sprintf("user-%s+%d", util.RandomString(5), i),
			Email: fmt.Sprintf("user+%d@example.com", i),
			// todo パスワードのハッシュを生成する
			Password: "test",
		})
	}

	bulk := make([]*ent.UserCreate, len(usersInputs))
	for i, input := range usersInputs {
		bulk[i] = tx.User.
			Create().
			SetID(input.ID).
			//SetOriginalID(input.OriginalID).
			SetEmail(input.Email).
			SetPassword(input.Password)
	}
	_, err := tx.User.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
