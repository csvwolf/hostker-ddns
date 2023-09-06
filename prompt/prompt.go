package prompt

import (
	"errors"
	ker "github.com/csvwolf/ker.go/model"
	"github.com/manifoldco/promptui"
	"net/mail"
)

type Prompt struct {
	Email  promptui.Prompt
	Token  promptui.Prompt
	Domain promptui.Prompt
	Record promptui.Select
}

func New() *Prompt {
	return &Prompt{}
}

func (*Prompt) SetEmail() (string, error) {
	prompt := promptui.Prompt{
		Label: "Hostker Email",
		Validate: func(s string) error {
			_, err := mail.ParseAddress(s)
			return err
		},
	}

	return prompt.Run()
}

func (*Prompt) SetToken() (string, error) {
	prompt := promptui.Prompt{
		Label: "Hostker Token",
		Validate: func(s string) error {
			if s == "" {
				return errors.New("token should not be empty")
			}
			return nil
		},
	}
	return prompt.Run()
}

func (*Prompt) SetDomain(domains []*ker.DomainListResponseItem) (string, error) {
	prompt := promptui.Select{
		Label: "Domain to set DDNS",
		Items: domains,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ .Domain }}",
			Active:   "\U0001F336 {{ .Domain | red }}",
			Inactive: "  {{ .Domain | cyan }}",
			Selected: "\U0001F336 {{ .Domain }}",
		},
	}
	i, _, err := prompt.Run()
	return domains[i].Domain, err
}

func (*Prompt) SetRecord(records []*ker.Record) (*ker.Record, error) {
	prompt := promptui.Select{
		Label: "Record to Update",
		Items: records,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ .Header }}({{ .Type }}): {{ .Value }}",
			Active:   "\U0001F336 {{ .Header | red }}({{ .Type | red }}): {{ .Value | red }}",
			Inactive: "  {{ .Header | cyan }}({{ .Type| cyan }}): {{ .Value | cyan }}",
			Selected: "\U0001F336 {{ .Header }}({{ .Type }}): {{ .Value }}",
		},
	}
	i, _, err := prompt.Run()
	return records[i], err
}
