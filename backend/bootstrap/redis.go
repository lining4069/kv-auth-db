package bootstrap

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/lining4069/ops-go/backend/global"
	"go.uber.org/zap"
	"strconv"
)

// InitializeRedis 初始化redis
func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.App.Config.Redis.Host + ":" + strconv.Itoa(global.App.Config.Redis.Port),
		Password: global.App.Config.Redis.Password,
		DB:       global.App.Config.Redis.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.App.Log.Error("Redis connect ping failed,err:", zap.Any("err", err))
		return nil
	}
	return client
}
