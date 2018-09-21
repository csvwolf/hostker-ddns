package ker

import (
	"encoding/json"
	"strconv"
)

// ListSshKey returns all sshKey
func (u *User) ListSshKey() (data ApiResponse){
	ret, err := request("listSshKey", nil, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}

// CreateSshKey creates a sshkey
func (u *User) CreateSshKey(name string, key string) (data ApiResponse){
	var b map[string]string
	b["name"] = name
	b["key"] = key
	ret, err := request("createSshKey", b, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}

// DeleteSshkey deletes a sshKey by ID
func (u *User) DeleteSshKey(id int) (data ApiResponse){
	var d map[string]string
	d["id"] = strconv.Itoa(id)
	ret, err := request("createSshKey", d, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}
