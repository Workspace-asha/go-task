package main

import (
 "log"
 "net/http"

 "go-task/internal/config"
 "go-task/internal/db"
 "go-task/internal/handlers"
 "go-task/internal/repository"
 "go-task/internal/service"
)

func main() {
 cfg := config.Load()

 pool := db.New(cfg.DatabaseURL)
 repo := repository.NewProjectRepo(pool)
 svc := service.NewProjectService(repo)
 h := handlers.New(svc)

 log.Println("Server started on :" + cfg.Port)
 http.ListenAndServe(":"+cfg.Port, h.Router())
}