package handler

import (
	"encoding/json"
	"fmt"
	"gotest/types"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

func (api *APIHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	usersList := api.service.GetAllUsers()
	json.NewEncoder(w).Encode(usersList)
}

func (api *APIHandler) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var user types.User
	json.Unmarshal(reqBody, &user)
	userCreated := api.service.CreateNewUser(user)
	json.NewEncoder(w).Encode(userCreated)
}

func (api *APIHandler) GetSingleUser(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	user, err := api.service.GetUserByName(name)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(user)
}

func (api *APIHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var updatedUser types.User
	json.Unmarshal(reqBody, &updatedUser)
	user, err := api.service.UpdateUser(updatedUser)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(user)
}

func (api *APIHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	_, err := api.service.DeleteUserByName(name)
	fmt.Println(err)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
