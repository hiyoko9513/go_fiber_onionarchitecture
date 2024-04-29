package configs

import (
	"github.com/gofiber/fiber/v2"
	"hiyoko-fiber/utils"
)

func NewServerConf() fiber.Config {
	return fiber.Config{
		DisableStartupMessage: !utils.Env("APP_DEBUG").GetBool(true),
	}
}
