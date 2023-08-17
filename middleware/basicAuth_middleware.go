package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

func hashString(str string) []byte {
	hash := sha256.New()
	hash.Write([]byte(str))
	return hash.Sum(nil)
}

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := hashString(username)
			passwordHash := hashString(password)
			expectedUsernameHash := hashString("test")
			expectedPasswordHash := hashString("root")

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash, expectedUsernameHash) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash, expectedPasswordHash) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
