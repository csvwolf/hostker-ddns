package command

import (
	"github.com/csvwolf/hostker-ddns/config"
	"github.com/csvwolf/hostker-ddns/prompt"
	"github.com/csvwolf/ker.go/api"
	"github.com/csvwolf/ker.go/constant"
	"github.com/csvwolf/ker.go/model"
	"github.com/spf13/cobra"
	"log"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "init project",
	Run: func(cmd *cobra.Command, arg []string) {
		var (
			existed bool
			err     error

			prompts         = prompt.New()
			conf            = config.New(nil)
			ker             *api.Ker
			domainList      *model.DomainListResponse
			records         []*model.Record
			recordsResponse *model.RecordsResponse
		)
		log.Println("Start to init config")
		// 新建配置目录
		if existed, err = conf.IfConfigIsExist(); err != nil {
			log.Fatalln(err)
			return
		} else if !existed {
			if err = conf.MkConfigDir(); err != nil {
				log.Fatalln(err)
				return
			}

			if err = conf.CreateConfig(); err != nil {
				log.Fatalln(err)
				return
			}
			log.Println("Create Config Success")
		}

		for conf.Email, err = prompts.SetEmail(); err != nil; {
			log.Println(err)
		}

		for conf.Token, err = prompts.SetToken(); err != nil; {
			log.Println(err)
		}

		ker = api.New(conf.Email, conf.Token)

		if domainList, err = ker.GetDomainList(); err != nil {
			log.Fatalln(err)
			return
		}

		for conf.Domain, err = prompts.SetDomain(domainList.Domains); err != nil; {
			log.Println(err)
		}

		if recordsResponse, err = ker.GetDomainRecords(conf.Domain); err != nil {
			log.Fatalln(err)
			return
		}

		for _, record := range recordsResponse.Records {
			if record.Type == constant.DomainTypeA || record.Type == constant.DomainTypeAAAA {
				records = append(records, record)
			}
		}

		for conf.Record, err = prompts.SetRecord(records); err != nil; {
			log.Println(err)
		}

		if err = conf.WriteConfig(); err != nil {
			log.Fatalln(err)
			return
		}

		log.Println("Save Config Success, pls use run command to test")
	},
}
