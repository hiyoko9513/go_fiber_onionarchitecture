package main

import (
	"fmt"

	"hiyoko-fiber/configs"
	"hiyoko-fiber/internal/infrastructure/database"
	"hiyoko-fiber/internal/interactor"
	"hiyoko-fiber/internal/presentation/http/app/middleware"
	"hiyoko-fiber/internal/presentation/http/app/router"
	"hiyoko-fiber/pkg/logging/file"
	"hiyoko-fiber/util"

	"github.com/gofiber/fiber/v2"
)

const (
	envRoot = "./cmd/app"
	logDir  = "./log/app"
)

var (
	databaseConf database.Conf
)

func init() {
	logger.SetLogDir(logDir)
	logger.Initialize()
	util.LoadEnv(envRoot)

	databaseConf = configs.NewMySqlConf()

	// todo timezone設定
}

func main() {
	f := fiber.New()
	entClient, err := database.NewMySqlConnect(databaseConf)
	if err != nil {
		logger.Fatal("failed to create dbclient", "error", err)
	}
	defer func(entClient *database.EntClient) {
		err := entClient.Close()
		if err != nil {
			logger.Fatal("failed to close dbclient", "error", err)
		}
	}(entClient)

	i := interactor.NewInteractor(entClient)
	h := i.NewAppHandler()

	router.NewRouter(f, h)
	middleware.NewMiddleware(f)
	if err := f.Listen(fmt.Sprintf(":%d", util.Env("SERVER_PORT").GetInt(8080))); err != nil {
		logger.Fatal("failed to start server", "error", err)
	}

	logger.Fatal(fmt.Sprintf("Server started on port: %d", util.Env("SERVER_PORT").GetInt(8080)))
}
