package repository

import (
	"context"
	"fmt"

	"hiyoko-fiber/internal/domain/services"
	"hiyoko-fiber/internal/infrastructure/database"
	"hiyoko-fiber/internal/pkg/ent/migrate"
)

// todo messageを細かく記載

type tableRepository struct {
	conn *database.EntClient
}

func NewTableRepository(conn *database.EntClient) services.TableRepository {
	return &tableRepository{conn}
}

func (r *tableRepository) Ping(ctx context.Context) error {
	err := r.conn.DB().PingContext(ctx)
	if err != nil {
		err = fmt.Errorf("failed to ping; error: %v", err)
		return err
	}
	return nil
}

func (r *tableRepository) Migrate(ctx context.Context) error {
	err := r.conn.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		err = fmt.Errorf("failed to migrate; error: %v", err)
		return err
	}
	return nil
}

func (r *tableRepository) TruncateAll(ctx context.Context) error {
	sqlclient := r.conn.DB()
	_, err := sqlclient.ExecContext(ctx, "SET FOREIGN_KEY_CHECKS=0;")
	if err != nil {
		err = fmt.Errorf("failed to truncate all tables; error: %v", err)
		return err
	}

	var truncateQuery string
	tables, err := sqlclient.QueryContext(ctx, "SELECT CONCAT('TRUNCATE TABLE ', GROUP_CONCAT(CONCAT('`',table_name,'`')),';') AS statement FROM information_schema.tables WHERE table_schema = 'hiyoko' AND table_name LIKE '%';")
	if err != nil {
		err = fmt.Errorf("failed to truncate all tables; error: %v", err)
		return err
	}
	tables.Next()
	err = tables.Scan(&truncateQuery)
	if err != nil {
		err = fmt.Errorf("failed to truncate all tables; error: %v", err)
		return err
	}
	err = tables.Close()
	if err != nil {
		err = fmt.Errorf("failed to truncate all tables; error: %v", err)
		return err
	}
	_, err = sqlclient.ExecContext(ctx, truncateQuery)
	if err != nil {
		err = fmt.Errorf("failed to truncate all tables; error: %v", err)
		return err
	}

	_, err = sqlclient.ExecContext(ctx, "SET FOREIGN_KEY_CHECKS=1;")
	if err != nil {
		err = fmt.Errorf("failed to truncate all tables; error: %v", err)
		return err
	}
	return nil
}

func (r *tableRepository) DropAll(ctx context.Context) error {
	sqlclient := r.conn.DB()
	_, err := sqlclient.ExecContext(ctx, "SET FOREIGN_KEY_CHECKS=0;")
	if err != nil {
		err = fmt.Errorf("failed to drop all tables; error: %v", err)
		return err
	}

	var truncateQuery string
	tables, err := sqlclient.QueryContext(ctx, "SELECT CONCAT('DROP TABLE ', GROUP_CONCAT(CONCAT('`',table_name,'`')),';') AS statement FROM information_schema.tables WHERE table_schema = 'hiyoko' AND table_name LIKE '%';")
	if err != nil {
		err = fmt.Errorf("failed to drop all tables; error: %v", err)
		return err
	}
	tables.Next()
	err = tables.Scan(&truncateQuery)
	if err != nil {
		err = fmt.Errorf("failed to drop all tables; error: %v", err)
		return err
	}
	err = tables.Close()
	if err != nil {
		return err
	}
	_, err = sqlclient.ExecContext(ctx, truncateQuery)
	if err != nil {
		err = fmt.Errorf("failed to drop all tables; error: %v", err)
		return err
	}

	_, err = sqlclient.ExecContext(ctx, "SET FOREIGN_KEY_CHECKS=1;")
	if err != nil {
		err = fmt.Errorf("failed to drop all tables; error: %v", err)
		return err
	}
	return nil
}
