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
	configure.AddConfigPath(os.Getenv("GOPATH") + "/src/github.com/csvwolf/hostker-ddns")
	configure.SetConfigName("config")
	err := configure.ReadInConfig()
	if err != nil {
		configure.WriteConfigAs("./config.json")
		log.Info("初始化配置文件中")
	}
}
