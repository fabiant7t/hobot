package server

import "fmt"

// ServerWrapper: API responses do wrap the servers.
type ServerWrapper struct {
	Server Server `json:"server"`
}

// Server: Representation used in list server responses.
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

func (srv *Server) String() string {
	name := srv.ServerName
	if name == "" {
		name = "[unnamed]"
	}
	return fmt.Sprintf("%-41s %-17s %-10s %-8s", name, srv.ServerIP, srv.DC, fmt.Sprintf("%d", srv.ServerNumber))
}

// DetailedServerWrapper: API responses do wrap the server.
type DetailedServerWrapper struct {
	Server DetailedServer `json:"server"`
}

func (srv *DetailedServer) String() string {
	name := srv.ServerName
	if name == "" {
		name = "[unnamed]"
	}
	return fmt.Sprintf("%-41s %-17s %-10s %-8s", name, srv.ServerIP, srv.DC, fmt.Sprintf("%d", srv.ServerNumber))
}

// DetailedServer: Reprsentation used in get server responses.
type DetailedServer struct {
	Server
	Reset            bool `json:"reset"`
	Rescue           bool `json:"rescue"`
	VNC              bool `json:"vnc"`
	Windows          bool `json:"windows"`
	Plesk            bool `json:"plesk"`
	CPanel           bool `json:"cpanel"`
	WOL              bool `json:"wol"`
	HotSwap          bool `json:"hot_swap"`
	LinkedStorageBox int  `json:"linked_storagebox"`
}

type Subnet struct {
	IP   string `json:"ip"`
	Mask string `json:"mask"`
}
