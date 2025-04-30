package server

type ServerListItem struct {
	Server Server `json:"server"`
}

type Server struct {
	ServerIP      string    `json:"server_ip"`
	ServerIPv6Net string    `json:"server_ipv6_net"`
	ServerNumber  int       `json:"server_number"`
	ServerName    string    `json:"server_name"`
	Product       string    `json:"product"`
	DC            string    `json:"dc"`
	Traffic       string    `json:"traffic"`
	Status        string    `json:"status"`
	Cancelled     bool      `json:"cancelled"`
	PaidUntil     string    `json:"paid_until"`
	IPs           []string  `json:"ip"`
	Subnets       *[]Subnet `json:"subnet"`
}

type Subnet struct {
	IP   string `json:"ip"`
	Mask string `json:"mask"`
}
