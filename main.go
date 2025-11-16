package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"prservice/db"
	"prservice/handlers"
	"prservice/repo"
	"prservice/service"

	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()

	pool, err := db.NewPostgresPool(ctx)
	if err != nil {
		log.Fatalf("db error: %v", err)
	}

	userRepo := repo.NewUserRepository(pool)
	teamRepo := repo.NewTeamRepository(pool)
	prRepo := repo.NewPRRepository(pool)

	userService := service.NewUserService(userRepo)
	teamService := service.NewTeamService(teamRepo, userRepo)
	prService := service.NewPRService(prRepo, userRepo, teamRepo)

	r := mux.NewRouter()

	handlers.RegisterUserRoutes(r, userService)
	handlers.RegisterTeamRoutes(r, teamService)
	handlers.RegisterPRRoutes(r, prService)

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	http.ListenAndServe(":"+port, r)
}
