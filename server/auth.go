package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dannyrsu/social-media-graphql/models"
	"github.com/dgrijalva/jwt-go"
)

func JwtValidate(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		requestPath := r.URL.Path
		authRoutes := []string{
			"/api/message/new",
		}

		for _, value := range authRoutes {
			if value == requestPath {
				tokenHeader := r.Header.Get("Authorization")
				if tokenHeader == "" {
					w.WriteHeader(http.StatusForbidden)
					WriteJsonResponse(w, map[string]interface{}{"error": "Missing Token"}, http.StatusForbidden)
					log.Fatalln("Missing token")
					return
				}

				splitToken := strings.Split(tokenHeader, " ")
				if len(splitToken) != 2 {
					w.WriteHeader(http.StatusForbidden)
					WriteJsonResponse(w, map[string]interface{}{"error": "Invalid or malformed token"}, http.StatusForbidden)
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
					WriteJsonResponse(w, map[string]interface{}{"error": "Invalid or malformed token"}, http.StatusForbidden)
					log.Fatalf("Invalid or malformed token: %v", err)
					return
				}

				ctx := context.WithValue(r.Context(), "user", tk.UserID)
				r = r.WithContext(ctx)
				h.ServeHTTP(w, r)
				return
			}
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
