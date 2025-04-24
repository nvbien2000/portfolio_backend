package http

import "github.com/gin-gonic/gin"

// SetUpRoutes is a function that sets up the routes for the API
func SetUpRoutes(router *gin.Engine) {

	router.GET("/health", HealthCheck)
	router.GET("/ping", PingHandler)

	// apiV1 := router.Group("/api/v1")
	// {
	// 	apiV1.GET("/blogs", blogHandler.GetAllBlogs) // Giả sử blogHandler được truyền vào SetupRoutes
	// 	apiV1.GET("/cv/download", cvHandler.DownloadCV) // Giả sử cvHandler được truyền vào SetupRoutes
	// }
}
