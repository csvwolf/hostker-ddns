package ip

import (
	"io"
	"net/http"
	"os"
)

// DefaultDetector 默认的探测用 ip，如果需要覆盖，请保证 body 中只返回 ip string
const DefaultDetector = "http://ip.codesky.me"
const EnvKey = "CUSTOM_IP_DETECTOR"

func Get() (string, error) {
	var (
		url  = os.Getenv(EnvKey)
		resp *http.Response
		body []byte
		err  error
	)

	// Get Environment Variable CUSTOM_IP
	if url == "" {
		url = DefaultDetector
	}

	if resp, err = http.Get(url); err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if body, err = io.ReadAll(resp.Body); err != nil {
		return "", err
	}

	return string(body), nil
}
