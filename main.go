package main

import (
	"os"

	"github.com/csvwolf/hostker-ddns/services"
	logging "github.com/op/go-logging"
	"github.com/urfave/cli"
)

var log = logging.MustGetLogger("ddns")

func main() {
	app := cli.NewApp()
	app.Name = "Hostker DDNS Register"
	app.Usage = "DDNS Register???"

	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "init the config of dns",
			Action: services.Init,
		},
		{
			Name:   "run",
			Usage:  "run to refresh ip",
			Action: services.Run,
		},
		{
			Name:   "start",
			Usage:  "register to crontab",
			Action: services.Start,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error(err.Error())
	}
}
