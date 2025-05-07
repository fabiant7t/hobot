package server

// ServerWrapper: API responses do wrap the servers.
type ServerWrapper struct {
	Server Server `json:"server"`
}

// Server: Representation used in list server responses.
type Server struct {
	ServerIP      string    `json:"server_ip" yaml:"server_ip"`
	ServerIPv6Net string    `json:"server_ipv6_net" yaml:"server_ipv6_net"`
	ServerNumber  int       `json:"server_number" yaml:"server_number"`
	ServerName    string    `json:"server_name" yaml:"server_name"`
	Product       string    `json:"product" yaml:"product"`
	DC            string    `json:"dc" yaml:"dc"`
	Traffic       string    `json:"traffic" yaml:"traffic"`
	Status        string    `json:"status" yaml:"status"`
	Cancelled     bool      `json:"cancelled" yaml:"cancelled"`
	PaidUntil     string    `json:"paid_until" yaml:"paid_until"`
	IPList        []string  `json:"ip" yaml:"ip"`
	SubnetList    *[]Subnet `json:"subnet" yaml:"subnet"`
}

// DetailedServerWrapper: API responses do wrap the server.
type DetailedServerWrapper struct {
	Server DetailedServer `json:"server"`
}

// DetailedServer: Reprsentation used in get server responses.
type DetailedServer struct {
	ServerIP         string    `json:"server_ip" yaml:"server_ip"`
	ServerIPv6Net    string    `json:"server_ipv6_net" yaml:"server_ipv6_net"`
	ServerNumber     int       `json:"server_number" yaml:"server_number"`
	ServerName       string    `json:"server_name" yaml:"server_name"`
	Product          string    `json:"product" yaml:"product"`
	DC               string    `json:"dc" yaml:"dc"`
	Traffic          string    `json:"traffic" yaml:"traffic"`
	Status           string    `json:"status" yaml:"status"`
	Cancelled        bool      `json:"cancelled" yaml:"cancelled"`
	PaidUntil        string    `json:"paid_until" yaml:"paid_until"`
	IPList           []string  `json:"ip" yaml:"ip"`
	SubnetList       *[]Subnet `json:"subnet" yaml:"subnet"`
	Reset            bool      `json:"reset" yaml:"reset"`
	Rescue           bool      `json:"rescue" yaml:"rescue"`
	VNC              bool      `json:"vnc" yaml:"vnc"`
	Windows          bool      `json:"windows" yaml:"windows"`
	Plesk            bool      `json:"plesk" yaml:"plesk"`
	CPanel           bool      `json:"cpanel" yaml:"cpanel"`
	WOL              bool      `json:"wol" yaml:"wol"`
	HotSwap          bool      `json:"hot_swap" yaml:"hot_swap"`
	LinkedStorageBox int       `json:"linked_storagebox" yaml:"linked_storagebox"`
}

type Subnet struct {
	IP   string `json:"ip" yaml:"ip"`
	Mask string `json:"mask" yaml:"mask"`
}
