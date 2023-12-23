package storage

import (
	"fmt"

	"github.com/luckyss6/cc_server/pkg/config"
	"github.com/luckyss6/cc_server/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Cfg.MysqlConfig.Username,
		config.Cfg.MysqlConfig.Password,
		config.Cfg.MysqlConfig.Host,
		config.Cfg.MysqlConfig.Port,
		config.Cfg.MysqlConfig.DBName,
	)

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("init db error: %s", err.Error()))
	}
	if config.Cfg.MysqlConfig.Debug {
		DB = DB.Debug()
	}
	// TODO: add db migrate
	autoMigrateDB()
}

func autoMigrateDB() {
	// 	haveUserInfo := DB.Migrator().HasTable(&model.UserInfo{})
	// 	if !haveUserInfo {
	// 		DB.Migrator().DropTable(&model.UserInfo{})
	// 		DB.Migrator().CreateTable(&model.UserInfo{})
	// 	}
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.UserInfo{})
}
