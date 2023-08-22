package handler

import (
	"encoding/json"
	"fmt"
	"gotest/jwt"
	"gotest/service"
	"gotest/types"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func AllGroceries(w http.ResponseWriter, r *http.Request) {
	var groceryList = service.GetAllGroceries()
	if len(groceryList) <= 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(groceryList)
}

func GetAllGroceriesFromUser(w http.ResponseWriter, r *http.Request) {
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
	userId, err := service.GetUserIdByUsername(username)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	groceryList, err := service.GetAllGroceriesFromUser(userId)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(groceryList)
}

func FindAllGroceriesByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	groceries := service.GetGroceryByName(name)
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
	_, err := service.CreateGrocery(&grocery)
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

	newGrocery, err := service.CreateGrocery(&grocery)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = service.CreateGroceryForUser(username, int(newGrocery.ID))
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
	err = service.DeleteGroceryForUser(username, convId)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteGrocery(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	convId, _ := strconv.Atoi(id)
	deletedGrocery, err := service.DeleteGroceryById(convId)
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
	newGrocery, err := service.UpdateGrocery(grocery)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newGrocery)
}

func UpdateStatusOfGrocery(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	status := r.FormValue("done")
	convId, _ := strconv.Atoi(id)
	convStatus, _ := strconv.ParseBool(status)
	newGrocery, err := service.UpdateStatusOfGrocery(convId, convStatus)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newGrocery)
}
