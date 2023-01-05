package routes

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"

	rep "github.com/shreyasprasad/reps/app/api/routes/rep"
	shared "github.com/shreyasprasad/reps/app/api/shared"
)

func SetupRoutes(class *shared.App) *gin.Engine {
	// Use Gin as router
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// Set routes for API
	// Update to POST, UPDATE, DELETE etc
	router.POST("/rep", rep.CreateRep)

	// Set up Gin Server
	return router
}
