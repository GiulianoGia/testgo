package handler

import (
	"encoding/json"
	"gotest/jwt"
	"gotest/service"
	"gotest/types"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type APIHandler struct {
	service service.Service
}

func NewAPIHandler(service service.Service) *APIHandler {
	return &APIHandler{
		service: service,
	}
}

func (api *APIHandler) AllGroceries(w http.ResponseWriter, r *http.Request) {
	var groceryList = api.service.GetAllGroceries()
	if len(groceryList) <= 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(groceryList)
}

func (api *APIHandler) GetAllGroceriesFromUser(w http.ResponseWriter, r *http.Request) {
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
	userId, err := api.service.GetUserIdByUsername(username)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	groceryList, err := api.service.GetAllGroceriesFromUser(userId)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(groceryList) <= 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(groceryList)
}

func (api *APIHandler) FindAllGroceriesByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	groceries := api.service.GetGroceryByName(name)
	if len(groceries) <= 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(groceries)
}

func (api *APIHandler) AddNewGrocery(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var grocery types.Grocery
	json.Unmarshal(reqBody, &grocery)
	_, err := api.service.CreateGrocery(&grocery)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (api *APIHandler) AddGroceryForUser(w http.ResponseWriter, r *http.Request) {
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

	newGrocery, err := api.service.CreateGrocery(&grocery)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = api.service.CreateGroceryForUser(username, int(newGrocery.ID))
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (api *APIHandler) DeleteGroceryFromUser(w http.ResponseWriter, r *http.Request) {
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
	err = api.service.DeleteGroceryForUser(username, convId)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (api *APIHandler) DeleteGrocery(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	convId, _ := strconv.Atoi(id)
	deletedGrocery, err := api.service.DeleteGroceryById(convId)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deletedGrocery)
}

func (api *APIHandler) UpadteGroceryById(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var grocery types.Grocery
	json.Unmarshal(reqBody, &grocery)
	newGrocery, err := api.service.UpdateGrocery(grocery)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newGrocery)
}

func (api *APIHandler) UpdateStatusOfGrocery(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	status := r.FormValue("done")
	convId, _ := strconv.Atoi(id)
	convStatus, _ := strconv.ParseBool(status)
	newGrocery, err := api.service.UpdateStatusOfGrocery(convId, convStatus)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newGrocery)
}
