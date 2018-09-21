package services

import (
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
	"github.com/violarulan/ker"
)

func Init(c *cli.Context) {
	prompt := promptui.Prompt{
		Label: "Hostker Email",
	}
	username, err := prompt.Run()

	if err != nil {
		log.Error(err.Error())
		return
	}

	prompt = promptui.Prompt{
		Label: "Hostker Token",
	}

	token, err := prompt.Run()

	if err != nil {
		log.Error(err.Error())
		return
	}

	prompt = promptui.Prompt{
		Label: "Domain to DDNS",
	}

	domain, err := prompt.Run()

	if err != nil {
		log.Error(err.Error())
		return
	}

	user := ker.User{Email: username, AccessToken: token}

	ok := user.Test()
	if !ok {
		log.Error("UnAuthorized")
		return
	}

	resp := user.GetDNSRecords(domain)

	filtered := []ker.DNSRecord{}

	for _, value := range resp.Records {
		if value.Type == "A" || value.Type == "AAAA" {
			filtered = append(filtered, value)
		}
	}

	promptSelect := promptui.Select{
		Label: "Select DNS Record to Change",
		Items: filtered,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ .Header }}({{ .Type }}): {{ .Data }}",
			Active:   "\U0001F336 {{ .Header | red }}({{ .Type | red }}): {{ .Data | red }}",
			Inactive: "  {{ .Header | cyan }}({{ .Type| cyan }}): {{ .Data | cyan }}",
			Selected: "\U0001F336 {{ .Header }}({{ .Type }}): {{ .Data }}",
		},
	}

	i, _, err := promptSelect.Run()

	if err != nil {
		log.Error(err.Error())
		return
	}

	recordId := filtered[i].ID

	configure.Set("email", username)
	configure.Set("token", token)
	configure.Set("domain", domain)
	configure.Set("id", recordId)

	configure.WriteConfig()
}
