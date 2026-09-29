package api

import "github.com/Wlczak/vlastas-vpn/state"

type ErrorResponse struct {
	Msg string
}

type GetLocationResponse struct {
	Location state.MullvadServerLocation
}
