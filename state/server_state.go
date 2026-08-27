package state

var state = &ServerState{}

func GetServerState() *ServerState {
	return state
}

func SetServerLocation(newLocation ServerGeoLocation) {
	state.CurrentLocation = newLocation
}
func SetServerLocationByCode(newLocationCode string) {
	for _, locationListItem := range state.LocationList {
		if locationListItem.Code == newLocationCode {
			state.CurrentLocation = locationListItem
		}
	}
}

func SetStateDefaults() {
	state.LocationList = []ServerGeoLocation{{Code: "cz", Name: "Czechia"}, {Code: "uk", Name: "United Kingdom"}, {Code: "jp", Name: "Japan"}}
	state.CurrentLocation = state.LocationList[0]
}
