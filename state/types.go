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
	Hostname             string
	CountryCode          string
	CountryName          string
	CityCode             string
	CityName             string
	Active               bool
	Owned                bool
	Provider             string
	IPv4AddrIn           string
	IPv6AddrIn           string
	NetworkPortSpeedGbit int
	Type                 string
	StatusMessages       []string
	Pubkey               string
	Daita                bool
}
