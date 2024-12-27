package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lining4069/ops-go/backend/app/common/request"
	"github.com/lining4069/ops-go/backend/controllers/app"
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
	// 测试注册
	router.POST("/user/register", func(c *gin.Context) {
		var form request.Register
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": request.GetErrorMsg(form, err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})

	})
	// 注册
	router.POST("/auth/register", app.Register)
	// 登录
	router.POST("/auth/login", app.Login)
}
