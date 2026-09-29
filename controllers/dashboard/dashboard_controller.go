package dashboard

import (
	"github.com/Wlczak/vlastas-vpn/state"
)

type AdminDashValues struct {
	Locations       []state.MullvadServerLocation
	CurrentLocation state.MullvadServerLocation
}
