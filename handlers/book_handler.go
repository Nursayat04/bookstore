package handlers

import (
	"Bookstore/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

var Books = make(map[int]models.Book)
var NextBookID = 1

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var list []models.Book

	for _, b := range Books {
		list = append(list, b)
	}

	json.NewEncoder(w).Encode(list)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var b models.Book

	json.NewDecoder(r.Body).Decode(&b)

	b.ID = NextBookID
	NextBookID++

	Books[b.ID] = b

	json.NewEncoder(w).Encode(b)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, _ := strconv.Atoi(idStr)

	b := Books[id]

	json.NewEncoder(w).Encode(b)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, _ := strconv.Atoi(idStr)

	var b models.Book
	json.NewDecoder(r.Body).Decode(&b)

	b.ID = id
	Books[id] = b

	json.NewEncoder(w).Encode(b)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, _ := strconv.Atoi(idStr)

	delete(Books, id)
}
