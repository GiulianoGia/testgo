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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	err := helper.DeleteUserByName(name)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
