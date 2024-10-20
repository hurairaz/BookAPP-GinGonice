package services

import (
	"errors"
	"gin-gonic/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (us *UserService) CreateUser(userReq *models.CreateUserRequest) (uint, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	newUser := models.User{Name: userReq.Name, Password: string(hashedPassword)}
	if err := us.db.Create(&newUser).Error; err != nil {
		return 0, err
	}
	return newUser.ID, nil
}

func (us *UserService) LoginUser(userReq *models.LoginUserRequest) (uint, error) {
	user := models.User{}
	if err := us.db.Where("name = ?", userReq.Name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("user not found")
		}
		return 0, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password)); err != nil {
		return 0, errors.New("incorrect password")
	}
	return user.ID, nil
}
