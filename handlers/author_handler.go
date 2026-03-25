package handlers

import (
	"Bookstore/models"
	"encoding/json"
	"net/http"
)

var Authors = make(map[int]models.Author)
var NextAuthorID = 1

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	var list []models.Author

	for _, a := range Authors {
		list = append(list, a)
	}

	json.NewEncoder(w).Encode(list)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var a models.Author

	json.NewDecoder(r.Body).Decode(&a)

	a.ID = NextAuthorID
	NextAuthorID++

	Authors[a.ID] = a

	json.NewEncoder(w).Encode(a)
}
