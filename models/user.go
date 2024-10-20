package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Password string
	Books    []Book `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
