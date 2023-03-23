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

	//flag.Parse()

	viper.SetConfigName(*conf)
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Read runtime config fail %s:", err.Error())
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	return nil
}
