package global

import (
	"github.com/go-redis/redis/v8"
	bitcask "github.com/lining4069/kv-auth-db/backend/app/bitcask"
	"github.com/lining4069/kv-auth-db/backend/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
	DB          *gorm.DB
	Redis       *redis.Client
	BitcaskDB   *bitcask.DB
}

var App = new(Application)
