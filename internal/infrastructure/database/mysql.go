package database

import (
	"database/sql"
	"fmt"
	"time"

	"hiyoko-fiber/internal/pkg/ent"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConf struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
	Debug    bool
}

type MysqlEntClient struct {
	*ent.Client
}

func NewMySqlConnect(conf MysqlConf) (*MysqlEntClient, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Name,
	)

	// sql.DB connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		err = fmt.Errorf("failed to connect to mysql; error: %v", err)
		return &MysqlEntClient{}, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	drv := entsql.OpenDB("mysql", db)

	client := ent.NewClient(ent.Driver(drv))

	if conf.Debug {
		client = client.Debug()
	}

	return &MysqlEntClient{client}, nil
}
