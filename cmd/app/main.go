package main

import (
	"fmt"

	"hiyoko-fiber/configs"
	"hiyoko-fiber/internal/infrastructure/database"
	"hiyoko-fiber/internal/interactor"
	"hiyoko-fiber/internal/presentation/http/app/middleware"
	"hiyoko-fiber/internal/presentation/http/app/router"
	logger "hiyoko-fiber/pkg/logging/file"
	"hiyoko-fiber/utils"

	"github.com/gofiber/fiber/v2"
)

const (
	envRoot = "./cmd/app"
	logDir  = "./log/app"
)

var (
	databaseConf database.MysqlConf
)

func init() {
	logger.SetLogDir(logDir)
	logger.Initialize()
	utils.LoadEnv(envRoot)

	databaseConf = configs.NewMySqlConf()

	utils.LoadTimezone(utils.Env("TZ").GetString())
}

func main() {
	f := fiber.New()
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

	i := interactor.NewInteractor(entClient)
	h := i.NewAppHandler()

	middleware.NewMiddleware(f)
	router.NewRouter(f, h)
	if err := f.Listen(fmt.Sprintf(":%d", utils.Env("SERVER_PORT").GetInt(8080))); err != nil {
		logger.Fatal("failed to start server", "error", err)
	}

	logger.Fatal(fmt.Sprintf("Server started on port: %d", utils.Env("SERVER_PORT").GetInt(8080)))
}
