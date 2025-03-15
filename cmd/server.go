package main

import (
	"event-planner/internal/db"
	"event-planner/internal/handlers"
	"event-planner/internal/models"
	"event-planner/internal/router"
	"event-planner/internal/services"
	"event-planner/pkg/auth"
	"event-planner/pkg/logger"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading env file", err)
	}

	env := os.Getenv("ENV")

	logger, err := logger.Init(env)
	if err != nil {
		log.Fatalln("Error initializing logger", err)
	}
	logger.Info("Logger initialized")

	db, err := db.Init()
	if err != nil {
		os.Exit(1)
	}
	defer db.Close()
	zap.S().Infow("Connected to postgres")

	model := models.New(db)
	auth := auth.New()
	service := services.New(model, auth)
	handler := handlers.New(service)

	router := router.New(handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5110"
	}

	zap.S().Info("Starting server on ", port)
	zap.S().DPanic(http.ListenAndServe(":"+port, router))
}
