package utils

import (
	"net/http"
	"strings"
	"sync"

	"github.com/dgrijalva/jwt-go"
)

var userMux sync.Mutex

func Authenticate(w http.ResponseWriter, r *http.Request, user interface{}) {
	name := r.FormValue("username")
	password := r.FormValue("password")

	if len(name) == 0 || len(password) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Username and password are required"))
		return
	}
	userMux.Lock()
	defer userMux.Unlock()
	token, err := getToken(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error generating JWT token " + err.Error()))
	} else {
		w.Header().Set("Authorization", "Bearer "+token)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Token: " + token))
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		authHeader = strings.Replace(authHeader, "Bearer", "", 1)
		claims, err := verifyToken(authHeader)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid token: " + err.Error()))
			return
		}
		name := claims.(jwt.MapClaims)["username"].(string)
		role := claims.(jwt.MapClaims)["role"].(string)
		r.Header.Set("username", name)
		r.Header.Set("role", role)
		if role != "admin" {
			w.WriteHeader(http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
