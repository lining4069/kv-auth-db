package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lining4069/ops-go/backend/app/middleware"
	"github.com/lining4069/ops-go/backend/app/services"
	"github.com/lining4069/ops-go/backend/controllers/app"
)

/*
存放api 分组路由
*/

// SetApiGroupRoutes 定义api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	// 注册
	router.POST("/auth/register", app.Register)
	// 登录
	router.POST("/auth/login", app.Login)
	// 使用Use给路由组增加中间件
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	// 获取当前登录用户信息
	authRouter.POST("/auth/info", app.Info)
}
