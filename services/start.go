package services

import "github.com/urfave/cli"

func Start(c *cli.Context) {
	err := Crontab.AddFunc("@every 1m", func() { Run(nil) })
	if err != nil {
		log.Info(err.Error())
	}
	Crontab.Start()
	select {} // Just Keep Alive
}
