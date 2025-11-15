package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"prservice/handlers"
	"prservice/repo"
	"prservice/service"

	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer pool.Close()

	repo := repo.NewPostgresRepo(pool)
	svc := service.NewService(repo)
	h := handlers.NewHandler(svc)

	r := mux.NewRouter()
	r.HandleFunc("/health", h.Health).Methods("GET")

	r.HandleFunc("/team/add", h.AddTeam).Methods("POST")
	r.HandleFunc("/team/get", h.GetTeam).Methods("GET")

	r.HandleFunc("/users/setIsActive", h.SetIsActive).Methods("POST")
	r.HandleFunc("/users/getReview", h.GetReviewsForUser).Methods("GET")

	r.HandleFunc("/pullRequest/create", h.CreatePR).Methods("POST")
	r.HandleFunc("/pullRequest/merge", h.MergePR).Methods("POST")
	r.HandleFunc("/pullRequest/reassign", h.ReassignReviewer).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("starting server on :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
