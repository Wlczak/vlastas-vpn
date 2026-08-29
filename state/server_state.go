package state

var state = &ServerState{}

func GetServerState() *ServerState {
	return state
}

func SetServerLocation(newLocation MullvadServerLocation) {
	state.CurrentLocation = newLocation
}
func SetServerLocationByHostname(newHostname string) {
	for _, mullvadLocationListItem := range state.LocationList {
		if mullvadLocationListItem.Hostname == newHostname {
			state.CurrentLocation = mullvadLocationListItem
		}
	}
}

func SetStateDefaults() {
	state.LocationList = []MullvadServerLocation{{Hostname: "default", CountryName: "DefaultLand"}}
	state.CurrentLocation = state.LocationList[0]
}
