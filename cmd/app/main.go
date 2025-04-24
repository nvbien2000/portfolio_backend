package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nvbien2000/portfolio_backend/internal/handler/http"
)

func main() {

	router := gin.Default()

	http.SetUpRoutes(router)

	router.Run(":8080")
}
