package main

import (
	"encoding/json"
	"fmt"
	"gotest/db"
	"gotest/helper"
	"gotest/types"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func AllGroceries(w http.ResponseWriter, r *http.Request) {
	var groceryList = helper.GetAllGroceries(db.DB)
	json.NewEncoder(w).Encode(groceryList)
}

func FindAllGroceriesByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	groceries := helper.GetGroceryByName(db.DB, name)
	json.NewEncoder(w).Encode(groceries)
}

func AddNewGrocery(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var grocery types.Grocery
	json.Unmarshal(reqBody, &grocery)
	helper.CreateGrocery(db.DB, &grocery)
}

func DeleteGrocery(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	convId, _ := strconv.Atoi(id)
	fmt.Println(convId)
	deletedGrocery := helper.DeleteGroceryById(db.DB, convId)
	json.NewEncoder(w).Encode(deletedGrocery)
}

func UpadteGroceryById(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var grocery types.Grocery
	json.Unmarshal(reqBody, &grocery)
	newGrocery := helper.UpdateGrocery(db.DB, grocery)
	json.NewEncoder(w).Encode(newGrocery)
}
