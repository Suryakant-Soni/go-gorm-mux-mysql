package controllers

import (
	"encoding/json"
	"go-gorm-mux-mysql/pkg/models"
	"go-gorm-mux-mysql/pkg/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var bookModel = models.Book{}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	allBooks, err := bookModel.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		res, _ := json.Marshal(allBooks)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookid"]
	intBookId, err := strconv.Atoi(bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	singleBook, _, err := bookModel.GetBookById(int64(intBookId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonBook, _ := json.Marshal(singleBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBook)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := models.Book{}
	utils.ParseBody(r, &createBook)
	b, err := createBook.CreateBook()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonB, err := json.Marshal(b)
	if err != nil {
		log.Fatalf("Error", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonB)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookid"]
	intId, _ := strconv.Atoi(bookId)
	deletedBook, err := bookModel.DeleteBook(int64(intId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resBook, _ := json.Marshal(deletedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdatedBook := models.Book{}
	utils.ParseBody(r, &UpdatedBook)
	params := mux.Vars(r)
	bookId := params["bookid"]
	intId, _ := strconv.Atoi(bookId)
	tempBook, db, err := bookModel.GetBookById(int64(intId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if UpdatedBook.Name != "" {
		tempBook.Name = UpdatedBook.Name
	}
	if UpdatedBook.Author != "" {
		tempBook.Author = UpdatedBook.Author
	}
	if UpdatedBook.Publication != "" {
		tempBook.Publication = UpdatedBook.Publication
	}
	db.Save(&tempBook)
	if db.Error != nil {
		http.Error(w, db.Error.Error(), http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(tempBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
