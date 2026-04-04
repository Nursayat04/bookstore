package handlers

import (
	"Bookstore/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var Books = make(map[int]models.Book)
var NextBookID = 1

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var bookList []models.Book

	for _, book := range Books {
		bookList = append(bookList, book)
	}

	json.NewEncoder(w).Encode(bookList)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	json.NewDecoder(r.Body).Decode(&book)

	book.ID = NextBookID
	NextBookID++

	Books[book.ID] = book

	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var updatedBook models.Book
	json.NewDecoder(r.Body).Decode(&updatedBook)

	updatedBook.ID = id
	Books[id] = updatedBook

	json.NewEncoder(w).Encode(updatedBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	delete(Books, id)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	book := Books[id]

	json.NewEncoder(w).Encode(book)
}
