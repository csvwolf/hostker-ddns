package ker

// ApiResponse includes all endpoints response types
// See each comment section for each endpoints
// For values, see http://www.hostker.com/about.html#/api
type ApiResponse struct {
	// @general
	Success      int64  `json:"success"`
	ErrorMessage string `json:errorMessage`

	// @quota
	Level         int64 `json:"level,string"`
	ServerCount   int64 `json:"servers,string"`
	HoldHours     int64 `json:"holdHours,string"`
	WillSuspend   int64 `json:"willSuspend,string"`
	Suspend       int64 `json:"suspend,string"`
	SuspendReason string `json:"suspendReason"`

	// @balance
	Balance int64 `json:"balance,string"`

	// @listSshKey
	Keys []SSHKey `json:"keys"`

	// @createSshKey
	Id string `json:"id,string"`

	// @deleteSshKey

	// @listServers
	Servers []Server `json:"servers"`

	// @getServer
	Server ServerDetails

	// @createServer
	UUID string `json:"uuid"`

	// @setPower

	// @setServer

	// @reinstallServer

	// @deleteServer

	// @cpuMonitor
	CPUData []CPUMonitorData `json:"data"`

	// @netMonitor
	NetData []NetMonitorData `json:"data"`

	// @natMonitor
	NatData []NetMonitorData `json:"data"`

	// @ioMonitor
	IOData []IOMonitorData `json:"data"`

	// @dnsGetRecords
	Records []DNSRecord `json:"records"`

	// @dnsAddRecord
	DNSRecordId int64 `json:"id"`

	// @dnsEditRecord

	// @dnsDeleteRecord
}

type SSHKey struct {
	Id      int64 `json:"id,string"`
	Name    string `json:"name"`
	Created int64 `json:"created,string"`
	SSHKey  string `json:"sshKey"`
}

type Server struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	MainIP  string `json:"mainip"`
	IntraIP string `json:"intraip"`
	Memory  int64 `json:"memory,string"`
	Area    string `json:"area"`
	Price   int64 `json:"price,string"`
}

type ServerDetails struct {
	Name             string `json:"name"`
	MainIP           string `json:"mainip"`
	IntraIP          string `json:"intraip"`
	NetMAC           string `json:"netMAC"`
	NatMAC           string `json:"natMAC"`
	Memory           int64 `json:"memory,string"`
	Area             string `json:"area"`
	IOBytes          int64 `json:"IOBytes,string"`
	IOPS             int64 `json:"IOPS,string"`
	OS               int64  `json:"OS,string"`
	NetIn            int64 `json:"netin,string"`
	NetOut           int64 `json:"netout,string"`
	NatIn            int64 `json:"natin,string"`
	NatOut           int64 `json:"natout,string"`
	ClosedDiskVirtIO int64 `json:"closedDiskVirtIO,string"`
	ClosedNetVirtIO  int64 `json:"closedNetVirtIO,string"`
	CPU              int64 `json:"cpu,string"`
	BootOrder        int64 `json:"bootOrder,string"`
	Price            int64 `json:"price,string"`
	ISO              int64 `json:"iso,string"`
	GateWay          string `json:"gateway"`
	NetMask          string `json:"netmask"`
	NatMask          string `json:"natmask"`
	RootPassword     string `json:"rootPassword"`
	HDDSize          int64 `json:"hddSize,string"`
	Created          int64 `json:"created,string"`
	Status           int64 `json:"status,string"`
	Power            int64 `json:"power,string"`
	SuspendReason    string `json:"suspendReason,string"`
}

type CPUMonitorData struct {
	TimeStamp int64
	CPULoad   int64
}

type NetMonitorData struct {
	TimeStamp int64 `json:"t"`
	BWIn      int64 `json:"bwin"`
	BWOut     int64 `json:"bwout"`
	TIn       int64 `json:"tin"`
	TOut      int64 `json:"tout"`
	PKGIn     int64 `json:"pkgin"`
	PKGOut    int64 `json:"pkgout"`
	PKGInS    float64 `json:"pkgins"`
	PKGOutS   float64 `json:"pkgouts"`
}

type IOMonitorData struct {
	TimeStamp int64 `json:"t"`
	IORqS     int64 `json:"iorqs"`
	IOWrS     float64 `json:"iowrs"`
	IORbS     int64 `json:"iorbs"`
	IOWbs     int64 `json:"iowbs"`
	IORq      int64 `json:"iorq"`
	IOWr      int64 `json:"iowr"`
	IOTr      int64 `json:"iotr"`
	IOTw      int64 `json:"iotw"`
}

type DNSRecord struct {
	ID       int64 `json:"id,string"`
	Header   string `json:"header"`
	Type     string `json:"type"`
	TTL      int64 `json:"ttl,string"`
	Priority int64 `json:"pritority,string"`
	Data     string `json:"data"`
}

/*
// Unmarshal twice to parse unknown complex json
func unmarshal(s []byte) (ApiResponse) {

	// Unmarshal twice and delete dumplicate keys
	resp := ApiResponse{}
	if err := json.Unmarshal(s, &resp); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(s, &resp.Data); err != nil {
		panic(err)
	}
	// delete keys
	delete(resp.Data, "success")
	delete(resp.Data, "errorMessage")
	return resp
}*/
