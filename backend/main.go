package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lining4069/ops-go/backend/bootstrap"
	"github.com/lining4069/ops-go/backend/global"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()
	// 初始化日志配置
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success !")

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			err := db.Close()
			if err != nil {
				global.App.Log.Error("err when close mysql", zap.Any("err", err))
			}
		}
	}()
	// 路由
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	// 启动服务器
	r.Run(":" + global.App.Config.App.Port)
}
