package state

var state = &ServerState{}

func GetServerState() *ServerState {
	return state
}

func SetServerLocation(newLocation ServerGeoLocation) {
	state.CurrentLocation = newLocation
}

func SetStateDefaults() {
	state.LocationList = []ServerGeoLocation{{Code: "location1", Name: "Location1"}, {Code: "location2", Name: "Location2"}, {Code: "location3", Name: "Location3"}}
	state.CurrentLocation = state.LocationList[0]
}
