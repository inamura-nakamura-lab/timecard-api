package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/inamura-nakamura-lab/timecard-api/domain/repository"
	"github.com/inamura-nakamura-lab/timecard-api/infrastructure/persistence/model"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	*gorm.DB
}

func NewUserRepository(orm *gorm.DB) repository.IUserRepository {
	return &userRepository{
		orm,
	}
}

func (repo *userRepository) InsertUser(ctx *gin.Context, user *model.User) error {
	return repo.DB.Create(user).Error
}

func (repo *userRepository) SelectUser(ctx *gin.Context, userID uint) (*model.User, error) {
	result := new(model.User)
	err := repo.DB.Where("id = ?", userID).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *userRepository) DeleteUser(ctx *gin.Context, userID uint) error {
	return repo.DB.Delete("id = ?", userID).Error
}


