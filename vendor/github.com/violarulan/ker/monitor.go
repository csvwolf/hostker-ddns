package ker

import (
	"encoding/json"
	"strconv"
)

// CPUMonitor retrives the cpu usage data in the given time
func (u *User) CPUMonitor(uuid string, startTime int, endTime int, unit string) (data ApiResponse) {
	b := make(map[string]string)
	b["uuid"] = uuid
	b["startTime"] = strconv.Itoa(startTime)
	b["endTime"] = strconv.Itoa(endTime)
	b["unit"] = unit
	ret, err := request("cpuMonitor", b, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}

// NetMonitor provide public external internet network statistics of specific virtual machine
// All outgoing internet traffic (e.g. website traffic)
func (u *User) NetMonitor(uuid string, startTime int, endTime int, unit string) (data ApiResponse) {
	b := make(map[string]string)
	b["uuid"] = uuid
	b["startTime"] = strconv.Itoa(startTime)
	b["endTime"] = strconv.Itoa(endTime)
	b["unit"] = unit
	ret, err := request("netMonitor", b, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}

// NatMonitor provide private internal network statistics of specific virtual machine
// All internal traffic are included (e.g. package source)
func (u *User) NatMonitor(uuid string, startTime int, endTime int, unit string) (data ApiResponse) {
	b := make(map[string]string)
	b["uuid"] = uuid
	b["startTime"] = strconv.Itoa(startTime)
	b["endTime"] = strconv.Itoa(endTime)
	b["unit"] = unit
	ret, err := request("natMonitor", b, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}

// IOMonitor gives IOPS data
func (u *User) IOMonitor(uuid string, startTime int, endTime int, unit string) (data ApiResponse) {
	b := make(map[string]string)
	b["uuid"] = uuid
	b["startTime"] = strconv.Itoa(startTime)
	b["endTime"] = strconv.Itoa(endTime)
	b["unit"] = unit
	ret, err := request("IOMonitor", b, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}
