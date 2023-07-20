package bootstrap

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func NewConfig() error {
	var (
		conf = flag.String("localConf", "./configs/runtime.config", "localConf file path")
	)

	// flag.Parse()

	// 设置配置文件的名字
	viper.SetConfigName(*conf)
	// 添加配置文件所在的路径,注意在 Linux 环境下 %GOPATH 要替换为 $GOPATH
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Read runtime config fail %s:", err.Error())
	}
	// viper.SetConfigType("yaml") // 设置配置文件类型

	// 监听配置文件的修改和变动
	viper.WatchConfig()
	// 监听回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	return nil
}
