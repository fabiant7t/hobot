package server

// ResetOptionWrapper: API responses do wrap the reset option.
type ResetOptionWrapper struct {
	ResetOption ResetOption `json:"reset"`
}

type ResetOption struct {
	ServerIP        string   `json:"server_ip" yaml:"server_ip"`
	ServerIPv6Net   string   `json:"server_ipv6_net" yaml:"server_ipv6_net"`
	ServerNumber    int      `json:"server_number" yaml:"server_number"`
	TypeList        []string `json:"type" yaml:"type"`
	OperatingStatus string   `json:"operating_status" yaml:"operating_status"`
}
