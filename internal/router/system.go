package router

import "github.com/gin-gonic/gin"

func SystemRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/ping", ping)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{"health": true})
}
