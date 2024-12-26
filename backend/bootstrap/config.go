package bootstrap

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/lining4069/ops-go/backend/global"
	"github.com/spf13/viper"
	"os"
)

func InitializeConfig() *viper.Viper {
	config := "config.yaml"

	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config faild %s \n", err))
	}

	// 箭筒配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changes", in.Name)
		// 重新配置
		if err := v.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})

	//将配置赋值给全局变量
	if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}
	return v
}
