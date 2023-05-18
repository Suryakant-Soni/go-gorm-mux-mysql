package models

import (
	"errors"
	"go-gorm-mux-mysql/pkg/config"

	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"column=bookname;type=varchar(100)" json:"name"`
	Author      string `gorm:"column=authorname;type=varchar(100)" json:"author"`
	Publication string `gorm:"column=publication;type=varchar(100)" json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
	if err := db.Create(&b).Error; err != nil {
		return nil, err
	} else {
		return b, nil
	}
}

func (b *Book) GetAllBooks() ([]Book, error) {
	Books := []Book{}
	err := db.Find(&Books).Error
	if err != nil {
		return nil, err
	} else {
		return Books, nil
	}
}

func (b *Book) GetBookById(Id int64) (*Book, *gorm.DB, error) {
	singleBook := Book{}
	db := db.Where("ID=?", Id).Find(&singleBook)
	if db.Error != nil {
		return nil, nil, db.Error
	} else {
		if db.RowsAffected == 0 {
			return nil, nil, errors.New("matching record doesmot exist")
		}
		return &singleBook, db, nil
	}
}

func (b *Book) DeleteBook(Id int64) (*Book, error) {
	book := Book{}
	err := db.Where("ID=?", Id).Delete(&book).Error
	if err != nil {
		return nil, err
	} else {
		return &book, nil
	}
}

// func (bookModel *Book) UpdateBook(b *Book) (*Book, error) {
// 	// update when the ID is given,
// 	book := Book{}
// 	// db.Where("ID=?", b.ID).Save(b)
// 	err := db.First(&book, b.ID).Save(&book).Error
// 	if err == nil {
// 		return &book, nil
// 	} else {
// 		return nil, err
// 	}
// }
