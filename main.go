package main

import (
	"github.com/csvwolf/hostker-ddns/command"
	"github.com/spf13/cobra"
	"log"
)

func main() {

	var (
		// https://cobra.dev/
		root = &cobra.Command{
			Use:   "hostker-ddns",
			Short: "auto detect and set to hostker dns platform",
			Long:  "...",
			Run:   func(cmd *cobra.Command, args []string) {},
		}
	)

	root.AddCommand(command.InitCmd)
	root.AddCommand(command.RunCmd)
	root.AddCommand(command.StartCmd)
	root.AddCommand(command.VersionCmd)
	if err := root.Execute(); err != nil {
		log.Fatalln(err)
		return
	}

}
