package mongo

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

// User => MongoDBに保存するやつ
type User struct {
	ID            bson.ObjectId `bson:"_id"`
	Uuid          string        `bson:"uuid"`
	Name          string        `bson:"name"`
	StudentNumber string        `bson:"student_num"`
	Date          time.Time     `bson:"date"`
}
