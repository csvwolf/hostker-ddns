package ker

import "encoding/json"

// Quota returns account quota
func (u *User) Quota() (data ApiResponse){
	ret, err := request("quota", nil, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}

// Balance gives account current balance
func (u *User) Balance() (data ApiResponse){
	ret, err := request("balance", nil, u)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(ret), &data)
	return data
}
