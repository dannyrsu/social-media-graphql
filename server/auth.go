package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dannyrsu/go-contacts/models"
	"github.com/dgrijalva/jwt-go"
)

func handleNoAuth(w http.ResponseWriter, r http.Request) {
	noAuth := []string{"/"}
	requestPath := r.URL.Path

	for _, value := range noAuth {
		if value == requestPath {
			h.ServeHTTP(w, r)
			return
		}
	}
}

func JwtValidate(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		
		handleNoAuth(w, r)
	
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			w.WriteHeader(http.StatusForbidden)
			WriteJsonMessage(w, "Missing Token")
			log.Fatalln("Missing token")
			return
		}

		splitToken := strings.Split(tokenHeader, " ")
		if len(splitToken) != 2 {
			w.WriteHeader(http.StatusForbidden)
			WriteJsonMessage(w, "Invalid or malformed token")
			log.Fatalln("Invalid or malformed token")
			return
		}

		tokenPart := splitToken[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusForbidden)
			WriteJsonMessage(w, "Invalid or malformed token")
			log.Fatalf("Invalid or malformed token: %v", err)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
