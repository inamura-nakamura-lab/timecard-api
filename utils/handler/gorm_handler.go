package handler

import (
	"github.com/inamura-nakamura-lab/timecard-api/infrastructure/persistence/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

func CreateConnection() *gorm.DB {
	VENDOR := os.Getenv("DB_VENDOR")
	USER     := os.Getenv("USER")
	PASS     := os.Getenv("PASS")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME   := os.Getenv("DBNAME")
	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+ "?parseTime=true"
	//CONNECT := "timecard"+":"+"timacard"+"@"+"tcp(127.0.0.1:3306)"+"/"+"timecard"+ "?parseTime=true"
	DB, err := gorm.Open(VENDOR, CONNECT)
	if err != nil {
		log.Println(err)
	}
	log.Println("[MIGRATION]: User Table")
	DB.AutoMigrate(&model.User{})
	return DB
}