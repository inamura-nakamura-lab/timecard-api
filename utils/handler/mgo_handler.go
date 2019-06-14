package handler

import (
	"github.com/globalsign/mgo"
	"log"
	"os"
)

var (
	hostName = os.Getenv("MONGO_HOST")
	userName = os.Getenv("MONGO_USER")
	password = os.Getenv("MONGO_DATABASE_PASSWORD")
	database = os.Getenv("MONGO_DATABASE")
	mongoURL = os.Getenv("MONGO_URL")
)

func CreateMgoConnection(collectionName string) *mgo.Collection {
	session, err := mgo.Dial(mongoURL)
	if err != nil {
		log.Println(err)
	}
	//defer session.Close()
	return session.DB(database).C(collectionName)
}
