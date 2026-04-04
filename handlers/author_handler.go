package handlers

import (
	"Bookstore/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

func GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	author := Authors[id]
	json.NewEncoder(w).Encode(author)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var a models.Author
	json.NewDecoder(r.Body).Decode(&a)

	a.ID = id
	Authors[id] = a

	json.NewEncoder(w).Encode(a)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	delete(Authors, id)
}
