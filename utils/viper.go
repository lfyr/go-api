package utils

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/lfyr/go-api/config"
	"github.com/spf13/viper"
	"os"
)

var (
	GVA_CONFIG config.Server
)

func init() {
	var conf string
	flag.StringVar(&conf, "c", "", "choose config file.")
	flag.Parse()
	if conf == "" { // 优先级: 命令行 > 环境变量 > 默认值
		if configEnv := os.Getenv("config_path"); configEnv == "" {
			conf = "./conf/config.yaml"
			fmt.Printf("您正在使用config的默认值,config的路径为%v\n", conf)
		} else {
			conf = configEnv
			fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", conf)
		}
	} else {
		fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", conf)
	}
	v := viper.New()
	v.SetConfigFile(conf)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	return
}
