package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/lining4069/ops-go/backend/global"
	"github.com/lining4069/ops-go/backend/routes"
	"go.uber.org/zap"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	//注册 api 分组路由
	apiGroup := router.Group("/api")
	routes.SetApiGroupRoutes(apiGroup)

	return router
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()
	err := r.Run(":" + global.App.Config.App.Port)
	if err != nil {
		global.App.Log.Error("server setup failed", zap.Any("err", err))
	}
}
