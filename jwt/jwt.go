package jwt

import (
	"errors"
	"gotest/types"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type JWTCLAIM struct {
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(username string, roleId int) (tokenString string, err error) {
	experationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTCLAIM{
		Username: username,
		Role:     roleId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: experationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return tokenString, err
}

func GetTokenFromRequestHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("Unauthorized")
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		return "", errors.New("Unauthorized")
	}

	return authHeaderParts[1], nil
}

func GetUsernameFromToken(token string) (username string, err error) {
	signKey, err := jwt.ParseWithClaims(
		token,
		&JWTCLAIM{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return "", err
	}
	claims, ok := signKey.Claims.(*JWTCLAIM)
	if !ok {
		err = errors.New("couldn't get claims")
		return "", err
	}
	return claims.Username, nil
}

func GetRoleFromToken(token string) (role string, err error) {
	signkey, err := jwt.ParseWithClaims(
		token,
		&JWTCLAIM{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return "", err
	}
	claims, ok := signkey.Claims.(*JWTCLAIM)
	if !ok {
		err := errors.New("couldn't get claims")
		return "", err
	}
	return types.GetRole(types.RoleEnum(claims.Role)), nil
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTCLAIM{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*JWTCLAIM)
	if !ok {
		err = errors.New("couldn't parse claims")
		return err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return err
	}
	return nil
}
