package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
存放api 分组路由
*/

// SetApiGroupRoutes 定义api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
