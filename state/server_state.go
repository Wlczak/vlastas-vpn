package state

var state = &ServerState{}

func GetServerState() *ServerState {
	return state
}

func SetStateDefaults() {
	locationService := state.LocationService
	locationService.SelectFirstAvailableLocation()
}
