package handlers

import (
	"Bookstore/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var Categories = make(map[int]models.Category)
var NextCategoryID = 1

func GetCategories(w http.ResponseWriter, r *http.Request) {
	var list []models.Category

	for _, c := range Categories {
		list = append(list, c)
	}

	json.NewEncoder(w).Encode(list)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var c models.Category
	json.NewDecoder(r.Body).Decode(&c)

	c.ID = NextCategoryID
	NextCategoryID++

	Categories[c.ID] = c

	json.NewEncoder(w).Encode(c)
}

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	category := Categories[id]
	json.NewEncoder(w).Encode(category)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var c models.Category
	json.NewDecoder(r.Body).Decode(&c)

	c.ID = id
	Categories[id] = c

	json.NewEncoder(w).Encode(c)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	delete(Categories, id)
}
