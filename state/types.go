package state

type ServerState struct {
	CurrentLocation MullvadServerLocation
	LocationList    []MullvadServerLocation
}

type ServerGeoLocation struct {
	Code string
	Name string
}

type MullvadServerLocation struct {
	Hostname             string   `json:"hostname"`
	CountryCode          string   `json:"country_code"`
	CountryName          string   `json:"country_name"`
	CityCode             string   `json:"city_code"`
	CityName             string   `json:"city_name"`
	Active               bool     `json:"active"`
	Owned                bool     `json:"owned"`
	Provider             string   `json:"provider"`
	IPv4AddrIn           string   `json:"ipv4_addr_in"`
	IPv6AddrIn           string   `json:"ipv6_addr_in"`
	NetworkPortSpeedGbit int      `json:"network_port_speed"`
	Type                 string   `json:"type"`
	StatusMessages       []string `json:"status_messages"`
	Pubkey               string   `json:"pubkey"`
	Daita                bool     `json:"daita"`
}
