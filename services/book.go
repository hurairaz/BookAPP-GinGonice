package services

import (
	"errors"
	"gin-gonic/models"
	"gorm.io/gorm"
)

type BookService struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{db: db}
}

func (bs *BookService) CreateBook(bookReq *models.CreateBookRequest) (*models.Book, error) {
	newBook := models.Book{Author: bookReq.Author, Title: bookReq.Title, UserID: bookReq.UserID}
	if err := bs.db.Create(&newBook).Error; err != nil {
		return nil, err
	}
	return &newBook, nil
}

func (bs *BookService) UpdateBook(bookReq *models.UpdateBookRequest, userID uint) (*models.Book, error) {
	book := models.Book{}
	if err := bs.db.First(&book, bookReq.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	if book.UserID != userID {
		return nil, errors.New("unauthorized: you do not have the permission to update this book")
	}
	if bookReq.Author != "" {
		book.Author = bookReq.Author
	}
	if bookReq.Title != "" {
		book.Title = bookReq.Title
	}
	if err := bs.db.Save(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (bs *BookService) DeleteBook(bookID uint, userID uint) error {
	book := models.Book{}
	if err := bs.db.First(&book, bookID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("book not found")
		}
		return err
	}
	if book.UserID != userID {
		return errors.New("unauthorized: you do not have the permission to delete this book")
	}
	if err := bs.db.Delete(&models.Book{}, bookID).Error; err != nil {
		return err
	}
	return nil
}

func (bs *BookService) GetBookByID(bookID uint) (*models.Book, error) {
	book := models.Book{}
	if err := bs.db.First(&book, bookID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return &book, nil
}

func (bs *BookService) GetBooksByAuthor(author string) ([]models.Book, error) {
	var books []models.Book
	if err := bs.db.Where("author = ?", author).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (bs *BookService) GetBooksByTitle(title string) ([]models.Book, error) {
	var books []models.Book
	if err := bs.db.Where("title = ?", title).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
