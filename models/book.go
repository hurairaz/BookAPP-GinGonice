package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string
	Author string
	UserID uint
}
type CreateBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	UserID uint
}
type UpdateBookRequest struct {
	ID     uint   `json:"id" binding:"required"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
