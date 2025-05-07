package server

type RescueOptionWrapper struct {
	RescueOption RescueOption `json:"rescue" yaml:"rescue"`
}

type RescueOption struct {
	ServerIP          string   `json:"server_ip" yaml:"server_ip"`
	ServerIPv6Net     string   `json:"server_ipv6_net" yaml:"server_ipv6_net"`
	ServerNumber      int      `json:"server_number" yaml:"server_number"`
	OS                []string `json:"os" yaml:"os"`
	Active            bool     `json:"active" yaml:"active"`
	Password          string   `json:"password" yaml:"password"`
	AuthorizedKeyList []string `json:"authorized_key" yaml:"authorized_key"`
	HostKeyList       []string `json:"host_key" yaml:"host_key"`
	BootTime          string   `json:"boot_time" yaml:"boot_time"`
}
