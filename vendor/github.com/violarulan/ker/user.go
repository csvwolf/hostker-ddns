// Ker, API wrapper for Hostker written in Go
// Author: @violarulan
// License: GPLv3
// https://github.com/violarulan/ker
package ker

import "encoding/json"

// User stores email and access token
type User struct {
	Email       string
	AccessToken string
}

// Test checks if User information are correct
// if correct returns True
func (u *User) Test() bool {
	var resp ApiResponse
	ret, err := request("quota", nil, u)
	if err != nil {
		return false
	}
	json.Unmarshal([]byte(ret), &resp)
	if int(resp.Success) != 1 {
		return false
	}
	return true
}