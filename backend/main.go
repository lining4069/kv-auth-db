package main

import (
	"github.com/lining4069/kv-auth-db/backend/bootstrap"
	"github.com/lining4069/kv-auth-db/backend/global"
	"go.uber.org/zap"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()

	// 初始化日志配置
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success !")

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()

	// 初始化Redis
	global.App.Redis = bootstrap.InitializeRedis()

	// 初始化bitcask kv 存储数据库
	global.App.BitcaskDB = bootstrap.InitialBitcaskDB()

	// 程序关闭前，释放mysql, bitcaskdb,redis 的占用
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			err := db.Close()
			if err != nil {
				global.App.Log.Error("err when close mysql", zap.Any("err", err))
			}
		}
		if global.App.BitcaskDB != nil {
			err := global.App.BitcaskDB.Close()
			if err != nil {
				global.App.Log.Error("err when close bitcask kv DB", zap.Any("err", err))
			}
		}

		if global.App.Redis != nil {
			err := global.App.Redis.Close()
			if err != nil {
				global.App.Log.Error("err when close redis ", zap.Any("err", err))
			}
		}
	}()

	// 初始化验证器
	bootstrap.InitializeValidator()

	// 启动服务器
	bootstrap.RunServer()
}
