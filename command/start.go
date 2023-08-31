package command

import (
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"log"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start ddns service",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Start ddns service, detecting every 1 minute")
		c := cron.New()
		if _, err := c.AddFunc("@every 1m", func() {
			run()
		}); err != nil {
			log.Fatalln(err)
			return
		}
		c.Start()

		select {}
	},
}
