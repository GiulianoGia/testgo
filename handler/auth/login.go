package handler

import (
	"encoding/json"
	"errors"
	"gotest/db"
	"gotest/helper"
	"gotest/jwt"
	"gotest/middleware"
	"gotest/types"
	"io"
	"net/http"
)

type AuthCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var userCredentials AuthCredentials
	json.Unmarshal(reqBody, &userCredentials)
	userCredentials.Password = string(middleware.HashString(userCredentials.Password))
	isAuthenticated, err := CheckUserCredentials(userCredentials.Username, userCredentials.Password)
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
	token, err := jwt.GenerateJWT(userCredentials.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(token)
	}
}

func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
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
	user, err := helper.GetUserByName(username)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func CheckAuthentication(w http.ResponseWriter, r *http.Request) {
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

func CheckUserCredentials(username string, password string) (authenticated bool, err error) {
	var user types.User
	err = db.DB.Where(&types.User{Name: username, Password: password}).Find(&user).Error
	if err != nil {
		return
	} else if user.Name == "" {
		err = errors.New("No user found")
		return
	}
	authenticated = true
	return
}
