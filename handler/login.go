package handler

import (
	"encoding/json"
	"errors"
	"gotest/jwt"
	"gotest/middleware"
	"io"
	"net/http"
)

type AuthCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (api *APIHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var userCredentials AuthCredentials
	json.Unmarshal(reqBody, &userCredentials)
	userCredentials.Password = string(middleware.HashString(userCredentials.Password))
	isAuthenticated, err := api.CheckUserCredentials(userCredentials.Username, userCredentials.Password)
	roleId, err := api.service.GetRoleIdByName(userCredentials.Username)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isAuthenticated {
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	token, err := jwt.GenerateJWT(userCredentials.Username, roleId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(token)
	}
}

func (api *APIHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
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
	user, err := api.service.GetUserByName(username)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (api *APIHandler) CheckAuthentication(w http.ResponseWriter, r *http.Request) {
	tokenString, err := jwt.GetTokenFromRequestHeader(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	err = jwt.ValidateToken(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (api *APIHandler) CheckUserCredentials(username string, password string) (authenticated bool, err error) {
	user, err := api.service.FindUserByUsernameAndPassword(username, password)
	if err != nil {
		return
	} else if user.Name == "" {
		err = errors.New("No user found")
		return
	}
	authenticated = true
	return
}
