package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/LordRahl90/little_quiz_backend/src/models"
	"github.com/LordRahl90/little_quiz_backend/src/utils"
	jwt "github.com/dgrijalva/jwt-go"
)

//JwtAuthentication This is a middleware that sits between the route and the operation.
var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		notAuth := []string{"/api/user/login", "/api/user/register"}
		requestPath := r.URL.Path
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		//if the path does not exist in the notAuth array, we proceed to verify the authenticity of the user.
		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") //retrieve the authorization header
		//This is if the authorization key is absent.
		if tokenHeader == "" {
			response = utils.Message(false, "Missing Authorization Key")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		splittedToken := strings.Split(tokenHeader, " ")
		if len(splittedToken) != 2 {
			response = utils.Message(false, "Invalid Token format.")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		tokenHalf := splittedToken[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenHalf, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response = utils.Message(false, "Malformed authentication Token")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		//Token is formed nicely but it is not valid.
		if !token.Valid {
			response = utils.Message(false, "Invalid Token")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		//At this point, all seems good.
		fmt.Printf("User %d\n", tk.UserID)
		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}
