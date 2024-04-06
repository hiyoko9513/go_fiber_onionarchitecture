package configs

import (
	"hiyoko-fiber/internal/infrastructure/database"
	"hiyoko-fiber/util"
)

func NewMySqlConf() (conf database.Conf) {
	conf.Host = util.Env("DB_HOST").GetString("localhost")
	conf.User = util.Env("DB_USER").GetString("hiyoko")
	conf.Password = util.Env("DB_PASSWORD").GetString("hiyoko")
	conf.Name = util.Env("DB_NAME").GetString("hiyoko")
	conf.Port = util.Env("DB_PORT").GetInt(3306)
	return
}
