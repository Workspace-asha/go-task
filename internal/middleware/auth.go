package middleware

import (
 "net/http"
 "strings"
 "go-task/internal/auth"
)

func JWTAuth(next http.Handler) http.Handler {
 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  authHeader := r.Header.Get("Authorization")
  if authHeader == "" {
   http.Error(w, "missing token", 401)
   return
  }

  tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
  _, err := auth.ValidateToken(tokenStr)
  if err != nil {
   http.Error(w, "invalid token", 401)
   return
  }

  next.ServeHTTP(w, r)
 })
}