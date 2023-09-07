package command

import (
	"github.com/csvwolf/hostker-ddns/config"
	"github.com/csvwolf/hostker-ddns/ip"
	"github.com/csvwolf/ker.go/api"
	"github.com/csvwolf/ker.go/model"
	"github.com/spf13/cobra"
	"log"
)

func run() error {
	var (
		conf      = config.New(nil)
		err       error
		ipAddr    string
		ker       *api.Ker
		newRecord *model.Record
	)
	if err = conf.GetConfig(); err != nil {
		return err
	}

	if ipAddr, err = ip.Get(); err != nil {
		return err
	}

	ker = api.New(conf.Email, conf.Token)
	if conf.Record.Value == ipAddr {
		log.Println("IP address is not changed")
		return nil
	}

	newRecord = &model.Record{
		Value:  ipAddr,
		Ttl:    conf.Record.Ttl,
		Type:   conf.Record.Type,
		Header: conf.Record.Header,
	}

	if _, err = ker.UpdateDomainRecord(conf.Domain, conf.Record, newRecord); err != nil {
		return err
	}
	conf.Record.Value = ipAddr
	if err = conf.WriteConfig(); err != nil {
		return err
	}
	log.Println("IP address updated:", ipAddr)
	return nil
}

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "run to test setting",
	Run: func(cmd *cobra.Command, args []string) {
		if err := run(); err != nil {
			log.Fatalln(err)
		}
	},
}
