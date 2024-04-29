package main

// exec command
// go run ./cmd/cli/db/main.go -query ping

import (
	"context"
	"flag"
	"path/filepath"

	"hiyoko-fiber/configs"
	"hiyoko-fiber/internal/infrastructure/database"
	"hiyoko-fiber/internal/interactor"
	logger "hiyoko-fiber/pkg/logging/file"
	"hiyoko-fiber/utils"
)

const (
	envRoot = "cmd/cli"
	logDir  = "./log/cli/db"

	dBQueryPing     = "ping"
	dBQueryMigrate  = "migrate"
	dBQuerySeed     = "seed"
	dBQueryTruncate = "truncate"
	dBQueryDrop     = "drop"

	errDefaultMsg = "Failed to query"
	successfulMsg = "Success query"
)

var (
	query *string
)

func init() {
	query = flag.String("query", "ping", "exec query")
	flag.Parse()

	logger.SetLogDir(logDir)
	logger.Initialize()
	logger.With("query", query)

	err := utils.EnvFile(filepath.Join(envRoot, ".env")).LoadEnv()
	if err != nil {
		logger.Fatal("Failed to load environment variables", "error", err)
	}
}

func main() {
	entClient, err := database.NewMySqlConnect(configs.NewMySqlConf())
	if err != nil {
		logger.Fatal("Failed to create dbclient", "error", err)
	}
	defer func(entClient *database.MysqlEntClient) {
		err := entClient.Close()
		if err != nil {
			logger.Fatal("Failed to close dbclient", "error", err)
		}
	}(entClient)

	ctx := context.Background()
	i := interactor.NewInteractor(entClient)
	r := i.NewTableRepository()

	switch *query {
	case dBQueryPing:
		err := r.Ping(ctx)
		if err != nil {
			logger.Fatal(errDefaultMsg, "error", err)
		}
	case dBQueryMigrate:
		err := r.Migrate(ctx)
		if err != nil {
			logger.Fatal(errDefaultMsg, "error", err)
		}
	case dBQuerySeed:
		err := r.Seed(ctx)
		if err != nil {
			logger.Fatal(errDefaultMsg, "error", err)
		}
	case dBQueryTruncate:
		err := r.TruncateAll(ctx)
		if err != nil {
			logger.Fatal(errDefaultMsg, "error", err)
		}
	case dBQueryDrop:
		err := r.DropAll(ctx)
		if err != nil {
			logger.Fatal(errDefaultMsg, "error", err)
		}
	}
	logger.Info(successfulMsg)
}
