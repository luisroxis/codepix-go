package db

import (
	"github.com/luisroxis/codepix-go/domain/model"
	"log"
	"os"
	"path/filepath"
	"runtime"
	
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	- "github.com//liq/pq"
	- "gorm.io/driver/sqlite"
)

func init() {
	_, b, _, _ := runtime.Caller(skip: 0)
	basepath := filepath.Dir(b)

	err:= godotenv..Load(basepath + "/../../.env")

	if err != nil {
		log.Fatalf(format: "Error loading .env file")
	}	
}

func ConnectDB(env string) *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	if env != "test" {
		dsn = os.Getenv(key: "dsn")
		db, err = gorm.Open(os.Getenv(key: "dbType"), dsn)
	} else {
		dsn = os.Getenv(key: "dsnTest")
		db, err = gorm.Open(os.Getenv(key: "dbTypeTest"), dsn)
	}

	if err != nil {
		log.Fatalf(format: "Error connect to database %s", err)
		panic(err)
	}

	if os.Getenv(key: "debug") == "true" {
		db.LogModel(enable: true)
	}

	if os.Getenv(key: "AutoMigrateDb") == "true" {
		db.AutoMigrate(&model.Bank{},&model.Account{}, &model.PixKey{}, &model.Transaction{}, &model.User{})
	}

	return db
}