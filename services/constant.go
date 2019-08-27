package services

import (
	"os"

	logging "github.com/op/go-logging"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

var configure = viper.New()
var log = logging.MustGetLogger("services")
var Crontab = cron.New()

func init() {
	path := os.Getenv("GOPATH") + "/src/github.com/csvwolf/hostker-ddns"
	configure.AddConfigPath(path)
	configure.SetConfigName("config")
	err := configure.ReadInConfig()
	if err != nil {
		f, err := os.Create(path + "/config.json")

		if err != nil {
			log.Error(err.Error())
		}

		defer f.Close()
		log.Info("初始化配置文件中")
	}
}
