package main

// exec command
// go run ./cmd/cli/db/main.go -query ping

import (
	"context"
	"flag"

	"hiyoko-fiber/configs"
	"hiyoko-fiber/internal/infrastructure/database"
	"hiyoko-fiber/internal/interactor"
	"hiyoko-fiber/pkg/logging/file"
	"hiyoko-fiber/utils"
)

const (
	EnvRoot = "cmd/cli"

	DBQueryPing     = "ping"
	DBQueryMigrate  = "migrate"
	DBQuerySeed     = "seed"
	DBQueryTruncate = "truncate"
	DBQueryDrop     = "drop"

	ErrDefaultMsg      = "failed to query"
	QuerySuccessfulMsg = "success query"
)

const logDir = "./log/cli/db"

var (
	databaseConf database.MysqlConf
	query        *string
)

func init() {
	// flag
	query = flag.String("query", "ping", "exec query")
	flag.Parse()

	logger.SetLogDir(logDir)
	logger.Initialize()
	logger.With("query", query)

	// load env
	utils.LoadEnv(EnvRoot)
	databaseConf = configs.NewMySqlConf()
}

func main() {
	entClient, err := database.NewMySqlConnect(databaseConf)
	if err != nil {
		logger.Fatal("failed to create dbclient", "error", err)
	}
	defer func(entClient *database.MysqlEntClient) {
		err := entClient.Close()
		if err != nil {
			logger.Fatal("failed to close dbclient", "error", err)
		}
	}(entClient)

	ctx := context.Background()
	i := interactor.NewInteractor(entClient)
	r := i.NewTableRepository()

	switch *query {
	case DBQueryPing:
		err := r.Ping(ctx)
		if err != nil {
			logger.Fatal(ErrDefaultMsg, "query", *query, "error", err)
		}
	case DBQueryMigrate:
		err := r.Migrate(ctx)
		if err != nil {
			logger.Fatal(ErrDefaultMsg, "query", *query, "error", err)
		}
	case DBQuerySeed:
		err := r.Seed(ctx)
		if err != nil {
			logger.Fatal(ErrDefaultMsg, "query", *query, "error", err)
		}
	case DBQueryTruncate:
		err := r.TruncateAll(ctx)
		if err != nil {
			logger.Fatal(ErrDefaultMsg, "query", *query, "error", err)
		}
	case DBQueryDrop:
		err := r.DropAll(ctx)
		if err != nil {
			logger.Fatal(ErrDefaultMsg, "query", *query, "error", err)
		}
	}
	logger.Info(QuerySuccessfulMsg)
}
