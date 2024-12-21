package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lai0xn/bsc-auth/pkg/utils"
)

func AuthMidleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "invalid token",
			})
			return
		}
		args := strings.Split(authHeader, " ")
		if len(args) < 2 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "please provide a token",
			})
			return

		}
		token, err := utils.ParseToken(args[1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "invalid token",
			})
			return

		}
		if !token.Valid {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "invalid token",
			})
			return

		}
		claims := token.Claims.(jwt.MapClaims)
		id, ok := claims["id"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "invalid token",
			})
			return

		}
		ctx := context.WithValue(r.Context(), "user", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
