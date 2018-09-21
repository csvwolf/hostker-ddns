package utils

import (
	"io/ioutil"
	"net/http"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("ip")

func GetIp(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	ip, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	return string(ip), nil
}
