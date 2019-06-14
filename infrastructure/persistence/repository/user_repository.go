package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/inamura-nakamura-lab/timecard-api/domain/repository"
	"github.com/inamura-nakamura-lab/timecard-api/infrastructure/persistence/model/mongo"
)

type userRepository struct{
	*mgo.Collection
}

func NewUserRepository(mgoConn *mgo.Collection) repository.IUserRepository {
	return &userRepository{mgoConn}
}

func (repo *userRepository) InsertUser(ctx *gin.Context, user *mongo.User) error {
	err := repo.Collection.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) SelectUsers(ctx *gin.Context) ([]*mongo.User, error) {
	var result []*mongo.User
	err := repo.Collection.Find(bson.M{}).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *userRepository) SelectUser(ctx *gin.Context, userID string) (*mongo.User, error) {
	var result *mongo.User
	err := repo.Collection.Find(bson.M{"uuid": userID}).One(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *userRepository) DeleteUser(ctx *gin.Context, userID string) error {
	err := repo.Collection.Remove(bson.M{"uuid": userID})
	if err != nil {
		return err
	}
	return nil
}
