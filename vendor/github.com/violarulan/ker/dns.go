package ker

import (
	"encoding/json"
	"strconv"
)

// GetDNSRecords returns all DNS records
func (u *User) GetDNSRecords(d string) (data ApiResponse){
	var b map[string]string
	b["domain"] = d
	ret, err := request("dnsGetRecords", b, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}

// DNSAddRecord adds a DNS record
func (u *User) DNSAddRecord(domain, header, dnstype, dnsdata string, ttl, priority int) (data ApiResponse){
	var b map[string]string
	b["domain"] = domain
	b["header"] = header
	b["type"] = dnstype
	b["data"] = dnsdata
	b["ttl"] = strconv.Itoa(ttl)
	b["priority"] = strconv.Itoa(priority)
	ret, err := request("dnsAddRecord", b, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}

// DNSEditRecord edit a specific DNS record by ID
func (u *User) DNSEditRecord(id int, dnsdata string, ttl int, priority int) (data ApiResponse){
	var b map[string]string
	b["data"] = dnsdata
	b["id"] = strconv.Itoa(id)
	b["ttl"] = strconv.Itoa(ttl)
	b["priority"] = strconv.Itoa(priority)
	ret, err := request("dnsEditRecord", b, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}

// DNSDeleteRecord deletes a DNS record by ID
func (u *User) DNSDeleteRecord(id int) (data ApiResponse){
	var b map[string]string
	b["id"] = strconv.Itoa(id)
	ret, err := request("dnsGetRecords", b, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}