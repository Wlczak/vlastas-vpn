package state

type ServerState struct {
	Location string
}

var state = &ServerState{}

func GetServerState() *ServerState {
	return state
}

func SetServerLocation(newLocation string) {
	state.Location = newLocation
}
