package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Server struct {
	System     System       `yaml:"system"`     // 系统配置
	Chaincodes []Chaincodes `yaml:"chaincodes"` // 链码证书配置
}

var config = new(Server)

func init() {
	v := viper.New()
	// v.SetConfigFile("../../config.yaml")
	v.SetConfigFile("config.yaml")

	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&config); err != nil {
			panic(err)
		}
	})
	if err := v.Unmarshal(&config); err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", config)
}

func Get() Server {
	return *config
}
