package server

// ResetWrapper: API responses do wrap the reset response.
type ResetWrapper struct {
	Reset Reset `json:"reset"`
}

type Reset struct {
	ServerIP      string `json:"server_ip" yaml:"server_ip"`
	ServerIPv6Net string `json:"server_ipv6_net" yaml:"server_ipv6_net"`
	ServerNumber  int    `json:"server_number" yaml:"server_number"`
	Type          string `json:"type" yaml:"type"`
}
