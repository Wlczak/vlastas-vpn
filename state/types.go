package state

type ServerState struct {
	CurrentLocation ServerGeoLocation
	LocationList    []ServerGeoLocation
}

type ServerGeoLocation struct {
	Code string
	Name string
}
