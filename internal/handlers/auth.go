package handlers

import (
 "encoding/json"
 "net/http"
 "go-task/internal/auth"
)

// Demo user
var demoUser = struct {
 ID       int
 Username string
 Password string
}{1, "admin", "password"}

func Login(w http.ResponseWriter, r *http.Request) {
 var req struct {
  Username string `json:"username"`
  Password string `json:"password"`
 }

 json.NewDecoder(r.Body).Decode(&req)

 if req.Username != demoUser.Username || req.Password != demoUser.Password {
  http.Error(w, "invalid credentials", 401)
  return
 }

 token, _ := auth.GenerateToken(demoUser.ID)
 json.NewEncoder(w).Encode(map[string]string{"token": token})
}