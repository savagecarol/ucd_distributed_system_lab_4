package apis

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(rg *gin.RouterGroup) {
	rg.GET("/first",first)
	rg.POST("/second",second)
}
