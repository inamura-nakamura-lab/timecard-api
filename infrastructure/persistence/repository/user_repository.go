package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/inamura-nakamura-lab/timecard-api/domain/repository"
	"github.com/inamura-nakamura-lab/timecard-api/infrastructure/persistence/model"
	"github.com/jinzhu/gorm"
	"time"
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

func (repo *userRepository) InsertAttendance(ctx *gin.Context, userID uint, dateFrom, dateTo string) error {
	var timeCards []model.TimeCard
	var timeCard model.TimeCard
	from, err := time.Parse("Thu May 24 22:56:30 JST 2001", dateFrom)
	if err != nil {
		return err
	}
	to, err := time.Parse("Thu May 24 22:56:30 JST 2001", dateTo)
	if err != nil {
		return err
	}
	err = repo.DB.Where("user_id = ?", userID).First(&timeCards).Error
	if err != nil {
		return err
	}
	if len(timeCards) == 0 {
		timeCard = model.TimeCard{
			UserID: userID,
		}
		attendance := model.Attendance{
			DateFrom:  from,
			DateTo: to,
		}
		timeCard.Attendances = append(timeCard.Attendances, attendance)
		err = repo.DB.Create(timeCard).Error
		if err != nil {
			return err
		}
	}else {
		var isExistNowTimeCard bool
		for _, value := range timeCards {
			year := time.Now().Year()
			month := time.Now().Month()
			timeCardYear := value.CreatedAt.Year()
			timeCardMonth := value.CreatedAt.Month()
			if year == timeCardYear && month == timeCardMonth {
				timeCard = value
				isExistNowTimeCard = true
				break
			}
			isExistNowTimeCard = false
		}
		if isExistNowTimeCard {
			attendance := model.Attendance{
				DateFrom:   from,
				DateTo:     to,
				TimeCardID: timeCard.ID,
			}
			timeCard.Attendances = append(timeCard.Attendances, attendance)
			err = repo.DB.Save(timeCard).Error
			if err != nil {
				return err
			}
		} else {
			timeCard = model.TimeCard{
				UserID: userID,
			}
			attendance := model.Attendance{
				DateFrom:  from,
				DateTo: to,
			}
			timeCard.Attendances = append(timeCard.Attendances, attendance)
			err = repo.DB.Create(timeCard).Error
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (repo *userRepository) SelectAttendance(ctx *gin.Context, userID uint) (*model.TimeCard, error) {
	var timeCards []model.TimeCard
	var timeCard model.TimeCard
	err := repo.DB.Where("user_id = ?", userID).First(&timeCards).Error
	if err != nil {
		return nil, err
	}
	for _, value := range timeCards {
		year := time.Now().Year()
		month := time.Now().Month()
		timeCardYear := value.CreatedAt.Year()
		timeCardMonth := value.CreatedAt.Month()
		if year == timeCardYear && month == timeCardMonth {
			timeCard = value
			break
		}
	}
	return &timeCard, nil
}