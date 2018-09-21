package ker

import (
	"encoding/json"
	"strconv"
	"strings"
)

// ListServers returns all server information
func (u *User) ListServers() (data ApiResponse){
	ret, err := request("listServers", nil, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}

// GetServer gives specific server information by UUID
func (u *User) GetServer(uuid string) (data ApiResponse){
	var b map[string]string
	b["uuid"] = uuid
	ret, err := request("getServer", b, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}

// CreateServer creates a new server
// Notice: key is sshKey ID, if not correct it will continue to set up new server without public key installed
func (u *User) CreateServer(name string, memory int, area string, os int, key string) (data ApiResponse){
	var b map[string]string
	b["name"] = name
	b["memory"] = strconv.Itoa(memory)
	b["area"] = strings.ToUpper(area)
	b["os"] = strconv.Itoa(os)
	b["sshKey"] = key
	ret, err := request("createServer", b, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}

// SetPower sets server power by UUID
// Notice: force shut off will lost all data in memory
// If a server is online, power on request will fail
func (u *User) SetPower(uuid, power string) (data ApiResponse){
	var b map[string]string
	b["uuid"] = uuid
	b["power"] = power
	ret, err := request("setPower", b, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}

// SetServer config a server by UUID
func (u *User) SetServer(uuid string, closeDiskVirtIO, closeNetVirtIO, isobootOrder int) (data ApiResponse){
	var b map[string]string
	b["uuid"] = uuid
	b["closeDiskVirtIO"] = strconv.Itoa(closeDiskVirtIO)
	b["closeNetVirtIO"] = strconv.Itoa(closeNetVirtIO)
	b["isobootOrder"] = strconv.Itoa(isobootOrder)
	ret, err := request("setServer", b, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}

// ReinstallServer reinstall a server OS by UUID
// Waring: All data is unrecoverable
func (u *User) ReinstallServer(uuid string, os int, key string) (data ApiResponse){
	var b map[string]string
	b["uuid"] = uuid
	b["os"] = strconv.Itoa(os)
	b["sshKey"] = key
	ret, err := request("reinstallServer", b, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}

// DeleteServer smashes a server OS by UUID
// Waring: All data is unrecoverable
func (u *User) DeleteServer(uuid string) (data ApiResponse){
	var d map[string]string
	d["uuid"] = uuid
	ret, err := request("deleteServer", d, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}