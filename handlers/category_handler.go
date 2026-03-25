package handlers

import (
	"Bookstore/models"
	"encoding/json"
	"net/http"
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
