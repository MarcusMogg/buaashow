package initialize

import (
	"buaashow/global"

	"go.uber.org/zap"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config"

func config() {
	v := viper.New()
	v.SetConfigName(defaultConfigFile)
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		zap.S().Fatalf("Fatal error config file: %s", err)
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Info("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GConfig); err != nil {
			zap.S().Error(err)
		}
	})

	if err := v.Unmarshal(&global.GConfig); err != nil {
		zap.S().Error(err)
	}
	zap.S().Debug(global.GConfig)
	global.GVP = v
}

func init() {
	loggerInit()
	config()
	resourceDir()
}
