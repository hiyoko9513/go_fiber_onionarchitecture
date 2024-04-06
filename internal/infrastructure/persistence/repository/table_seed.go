package repository

import (
	"context"
	"fmt"

	"hiyoko-fiber/internal/infrastructure/persistence/repository/seeds"
)

const (
	defaultErrMsg = "failed to seed; error: %v"
)

func (r *tableRepository) Seed(ctx context.Context) error {
	tx, err := r.conn.Tx(ctx)
	if err != nil {
		err = fmt.Errorf(defaultErrMsg, err)
		return err
	}

	err = seeds.UsersSeed(ctx, tx)
	if err != nil {
		err = rollback(tx, err)
		err = fmt.Errorf(defaultErrMsg, err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		err = rollback(tx, err)
		err = fmt.Errorf(defaultErrMsg, err)
		return err
	}
	return nil
}
