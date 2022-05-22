package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/t0239184/golearn/internal/database"
	"github.com/t0239184/golearn/internal/middleware"
)

func New(db *database.GormDatabase) *gin.Engine {
	router := gin.New()
	router.HandleMethodNotAllowed = true

	/* CORS */
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	middleware.Setup()
	/* Common Middleware */
	router.Use(
		// gin.Logger(),
		middleware.LoggerToFile(),
		gin.Recovery(),
		// middleware.ErrorHandler(),
		middleware.Trace(),
		middleware.ReqResAopLogging(),
		middleware.ResponseHeader(),
		cors.New(config),
	)

	/* API Mapping */
	sysApiGroup := router.Group("/sys")
	SystemRouter(sysApiGroup)

	v1ApiGroup := router.Group("/api/v1")
	UserRouter(v1ApiGroup, db)

	return router
}
