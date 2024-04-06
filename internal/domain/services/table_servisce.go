package services

import (
	"context"
)

type TableRepository interface {
	Ping(ctx context.Context) error
	Migrate(ctx context.Context) error
	Seed(ctx context.Context) error
	TruncateAll(ctx context.Context) error
	DropAll(ctx context.Context) error
}
