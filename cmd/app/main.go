package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nvbien2000/portfolio_backend/internal/config"
	"github.com/nvbien2000/portfolio_backend/internal/handler/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	router := gin.Default()

	http.SetUpRoutes(router)

	serverAddress := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Starting server on %s", serverAddress)

	if err := router.Run(serverAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
