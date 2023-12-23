package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gofiber/storage/mysql"
	"github.com/luckyss6/cc_server/pkg/config"
)

var DB *sql.DB

func InitDB() {
	connURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Cfg.MysqlConfig.Username,
		config.Cfg.MysqlConfig.Password,
		config.Cfg.MysqlConfig.Host,
		config.Cfg.MysqlConfig.Port,
		config.Cfg.MysqlConfig.DBName,
	)

	storage := mysql.New(mysql.Config{
		ConnectionURI: connURL,
		Reset:         false,
		GCInterval:    10 * time.Second,
	})

	err := storage.Conn().Ping()
	if err != nil {
		panic("failed to connect database")
	}

	DB = storage.Conn()
}
