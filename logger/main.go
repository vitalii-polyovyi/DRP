package main

import (
	"drp/logger/database"
	"drp/logger/middlewares"
	"drp/logger/routes"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func startServer(router *chi.Mux) {
	port := ":" + os.Getenv("APP_PORT")
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}

	fmt.Println("Logging API Service running @ http://127.0.0.1" + port)

	if err := http.Serve(l, router); err != nil {
		database.CloseConnection()
		if errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server closed\n")
		} else if err != nil {
			log.Fatalf("Serve error: %s\n", err)
		}
	}
}

func main() {
	loadEnv()

	database.GetClient()
	defer database.CloseConnection()

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.SetHeader("Content-Type", "application/json"))
	router.Use(middleware.Compress(6))
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	})
	router.Use(cors.Handler)
	router.Use(middlewares.AuthMiddleware)

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	router.Mount("/http-log", routes.HttpLogRoutes{}.Routes())
	router.Mount("/entity-log", routes.EntityLogRoutes{}.Routes())
	router.Mount("/event-log", routes.EventLogRoutes{}.Routes())

	startServer(router)
}
