package handler

import (
	"fmt"
	"github.com/inamura-nakamura-lab/timecard-api/infrastructure/persistence/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var (
	// GormConn is Gorm Connection
	GormConn *gorm.DB
)

func init() {
	GormConn = CreateConnection()
}

func CreateConnection() *gorm.DB {
	VENDOR := os.Getenv("DB_VENDOR")
	USER     := os.Getenv("USER")
	PASS     := os.Getenv("PASS")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME   := os.Getenv("DBNAME")
	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+ "?parseTime=true"
	DB, err := gorm.Open(VENDOR, CONNECT)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Create User Table")
	DB.AutoMigrate(&model.User{})
	return DB
}