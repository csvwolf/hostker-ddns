package config

import (
	"errors"
	"github.com/csvwolf/ker.go/constant"
	ker "github.com/csvwolf/ker.go/model"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

const DefaultFilename = "ddns_config"

type Options struct {
	ConfigDir string
	FileName  string
}

type Config struct {
	Email  string
	Domain string
	Token  string
	Record *ker.Record

	v          *viper.Viper
	configDir  string
	configName string
}

func defaultDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "./.hostker"
	}
	return dirname + "/.hostker"
}

func interfaceToRecord(item map[string]interface{}) *ker.Record {
	return &ker.Record{
		Type:   constant.DomainType(item["type"].(string)),
		Ttl:    uint32(item["ttl"].(int)),
		Value:  item["value"].(string),
		Header: item["header"].(string),
	}
}

func New(options *Options) *Config {
	var (
		dir  = defaultDir()
		conf = &Config{
			configDir:  dir,
			configName: DefaultFilename,
			v:          viper.New(),
		}
	)

	conf.v.SetConfigType("yaml")

	if options == nil {
		conf.v.AddConfigPath(conf.configDir)
		conf.v.SetConfigName(conf.configName)
		return conf
	}

	if options.ConfigDir != "" {
		conf.configDir = options.ConfigDir
	}
	if options.FileName != "" {
		conf.configName = options.FileName
	}

	conf.v.AddConfigPath(conf.configDir)
	conf.v.SetConfigName(conf.configName)
	return conf
}

func (config *Config) CreateConfig() error {
	_, err := os.Create(filepath.Join(config.configDir, config.configName+".yaml"))
	return err
}

func (config *Config) IfConfigIsExist() (bool, error) {
	var (
		err                     error
		configFileNotFoundError viper.ConfigFileNotFoundError
	)
	if err = config.v.ReadInConfig(); err != nil {
		if errors.As(err, &configFileNotFoundError) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (config *Config) MkConfigDir() error {
	var err error
	if _, err = os.Stat(config.configDir); os.IsNotExist(err) {
		err = os.Mkdir(config.configDir, os.ModePerm)
	}
	return err
}

func (config *Config) SetEmail(email string) {
	config.v.Set("email", email)
}

func (config *Config) SetToken(token string) {
	config.v.Set("token", token)
}

func (config *Config) SetDomain(domain string) {
	config.v.Set("domain", domain)
}

func (config *Config) SetRecord(record *ker.Record) {
	config.v.Set("record", record)
}

func (config *Config) GetConfig() error {
	if err := config.v.ReadInConfig(); err != nil {
		return err
	}
	config.Domain = config.v.Get("domain").(string)
	config.Email = config.v.Get("email").(string)
	config.Token = config.v.Get("token").(string)
	config.Record = interfaceToRecord(config.v.Get("record").(map[string]interface{}))

	return nil
}

func (config *Config) WriteConfig() error {
	config.SetDomain(config.Domain)
	config.SetToken(config.Token)
	config.SetEmail(config.Email)
	config.SetRecord(config.Record)

	if err := config.v.WriteConfig(); err != nil {
		return err
	}
	return nil
}
