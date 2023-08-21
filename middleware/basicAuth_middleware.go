package middleware

import (
	"crypto/sha256"
	"encoding/base64"
	"gotest/jwt"
	"net/http"
)

func HashString(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := HashString(username)
			passwordHash := HashString(password)
			expectedUsernameHash := HashString("test")
			expectedPasswordHash := HashString("root")

			usernameMatch := usernameHash == expectedUsernameHash
			passwordMatch := passwordHash == expectedPasswordHash

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := jwt.GetTokenFromRequestHeader(r)
		if err != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		err = jwt.ValidateToken(token)
		if err != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		} else {
			next.ServeHTTP(w, r)
			return
		}
	})
}
