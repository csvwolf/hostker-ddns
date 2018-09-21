package ker

import (
	"net/http"
	"net/url"
	"bytes"
	"io/ioutil"
)

// request implants a basic HTTP client
func request(endpoint string, data map[string]string, u *User) (ret string, err error){
	body := encode(data)
	body.Set("email", u.Email)
	body.Set("token", u.AccessToken)
	url := BASE_URL + "/" + endpoint
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(body.Encode()))
	if err != nil {
		return ret, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Client", "Ker@" + VERSION)
	resp, err := client.Do(req)
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	ret = string(r)
	return

}

// encode converts a map to url.Values
func encode(data map[string]string) url.Values {
	ret := url.Values{}
	for k, v := range data {
		ret.Set(k, v)
	}
	return ret
}