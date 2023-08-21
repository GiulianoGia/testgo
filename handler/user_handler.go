package handler

import (
	"encoding/json"
	"gotest/helper"
	"gotest/types"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	usersList := helper.GetAllUsers()
	json.NewEncoder(w).Encode(usersList)
}

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var user types.User
	json.Unmarshal(reqBody, &user)
	userCreated := helper.CreateNewUser(user)
	json.NewEncoder(w).Encode(userCreated)
}

func GetSingleUser(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	user, err := helper.GetUserByName(name)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var updatedUser types.User
	json.Unmarshal(reqBody, &updatedUser)
	user, err := helper.UpdateUser(updatedUser)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	_, err := helper.DeleteUserByName(name)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
