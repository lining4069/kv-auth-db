package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lining4069/kv-auth-db/backend/app/middleware"
	"github.com/lining4069/kv-auth-db/backend/app/services"
	"github.com/lining4069/kv-auth-db/backend/controllers/app"
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
	{
		// 获取当前登录用户信息
		authRouter.POST("/auth/info", app.Info)
		// 登出
		authRouter.POST("/auth/logout", app.Logout)
	}
	bitcaskRouter := router.Group("bitcask").Use(middleware.JWTAuth(services.AppGuardName))
	{
		bitcaskRouter.POST("/put", app.BitcaskPut) // bitcask 数据库put
		bitcaskRouter.GET("/get", app.BitcaskGet)  // bitcask 数据库put
	}

}
