package router

import (
	"github.com/gin-gonic/gin"
	"github.com/t0239184/golearn/internal/database"
	v1 "github.com/t0239184/golearn/internal/router/api/v1"
)

func UserRouter(routerGroup *gin.RouterGroup, db *database.GormDatabase) {
	v1 := v1.NewUserHandler(db)

	user := routerGroup.Group("/user")
	user.GET("/", v1.QueryAllUser)
	user.GET("/:id", v1.FindUserById)
	user.POST("/", v1.CreateUser)
	user.POST("/:id", v1.UpdateUser)
	user.DELETE("/:id", v1.DeleteUser)
}
