package handler

import (
	"encoding/json"
	"fmt"
	"gotest/helper"
	"gotest/jwt"
	"gotest/types"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func AllGroceries(w http.ResponseWriter, r *http.Request) {
	var groceryList = helper.GetAllGroceries()
	if len(groceryList) <= 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(groceryList)
}

func FindAllGroceriesByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	groceries := helper.GetGroceryByName(name)
	if len(groceries) <= 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(groceries)
}

func AddNewGrocery(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var grocery types.Grocery
	json.Unmarshal(reqBody, &grocery)
	_, err := helper.CreateGrocery(&grocery)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func AddGroceryForUser(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.GetTokenFromRequestHeader(r)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	username, err := jwt.GetUsernameFromToken(token)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	reqBody, _ := io.ReadAll(r.Body)
	var grocery types.Grocery
	json.Unmarshal(reqBody, &grocery)

	newGrocery, err := helper.CreateGrocery(&grocery)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = helper.CreateGroceryForUser(username, int(newGrocery.ID))
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func DeleteGroceryFromUser(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.GetTokenFromRequestHeader(r)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	username, err := jwt.GetUsernameFromToken(token)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id := chi.URLParam(r, "id")
	convId, _ := strconv.Atoi(id)
	fmt.Println(username, convId)
	err = helper.DeleteGroceryForUser(username, convId)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteGrocery(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	convId, _ := strconv.Atoi(id)
	deletedGrocery, err := helper.DeleteGroceryById(convId)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deletedGrocery)
}

func UpadteGroceryById(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var grocery types.Grocery
	json.Unmarshal(reqBody, &grocery)
	newGrocery, err := helper.UpdateGrocery(grocery)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newGrocery)
}
