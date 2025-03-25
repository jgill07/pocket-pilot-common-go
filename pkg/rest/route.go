package rest

import "github.com/gin-gonic/gin"

func SetupRouter(healthCheck gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.ContextWithFallback = true

	if healthCheck != nil {
		router.GET("/healthz", healthCheck)
	}

	return router
}
