
package main

import (
 "context"
 "log"
 "net/http"
 "os/signal"
 "syscall"

 "task-api/internal/config"
 "task-api/internal/db"
 "task-api/internal/handlers"
 "task-api/internal/repository"
 "task-api/internal/service"
)

func main() {
 cfg := config.Load()

 pool := db.New(cfg.DatabaseURL)
 repo := repository.NewProjectRepo(pool)
 svc := service.NewProjectService(repo)
 handler := handlers.New(svc)
 srv := &http.Server{
	Addr:    ":" + cfg.Port,
	Handler: handler.Router(),
   }
  
   go func() {
	log.Println("server started")
	if err := srv.ListenAndServe(); err != nil {
	 log.Fatal(err)
	}
   }()
  
   ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
   defer stop()
   <-ctx.Done()
   srv.Shutdown(context.Background())
  }