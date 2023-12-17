package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	dataSourceName := "root:5201314123Huang!@tcp(127.0.0.1:3306)/?parseTime=True"
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
		os.Exit(-1)
	}

	testDB = db.Debug()
	m.Run()
}

func TestInitDB(t *testing.T) {
	
}
