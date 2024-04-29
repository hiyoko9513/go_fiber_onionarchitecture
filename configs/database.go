package configs

import (
	"hiyoko-fiber/internal/infrastructure/database"
	"hiyoko-fiber/utils"
)

func NewMySqlConf() (conf database.MysqlConf) {
	conf.Host = utils.Env("DB_HOST").GetString()
	conf.User = utils.Env("DB_USER").GetString()
	conf.Password = utils.Env("DB_PASSWORD").GetString()
	conf.Name = utils.Env("DB_NAME").GetString()
	conf.Port = utils.Env("DB_PORT").GetInt()
	conf.Debug = utils.Env("APP_DEBUG").GetBool(false)
	return
}
