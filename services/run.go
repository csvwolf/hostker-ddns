package services

import (
	"github.com/csvwolf/hostker-ddns/utils"
	"github.com/urfave/cli"
	"github.com/violarulan/ker"
)

func Run(c *cli.Context) {
	ip, err := utils.GetIp("http://ip.codesky.me")
	log.Info(ip)
	if err != nil {
		return
	}
	user := ker.User{Email: configure.GetString("email"), AccessToken: configure.GetString("token")}
	ok := user.Test()
	if !ok {
		log.Error("UnAuthorized")
		return
	}
	if configure.GetString("ip") == ip {
		log.Info("The Same value")
		return
	}
	resp := user.DNSEditRecord(configure.GetInt("id"), ip, 300, 0)
	if resp.Success != 1 {
		log.Error(resp.ErrorMessage)
	} else {
		configure.Set("ip", ip)
		configure.WriteConfig()
		log.Info("Success Change!")
	}
}
