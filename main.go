package main

import (
	h "github.com/Wlczak/vlastas-vpn/helpers"
	"github.com/Wlczak/vlastas-vpn/state"
)

func main() {
	r := SetupRouter()

	go func() {
		err := r.Run()
		h.PanicChckErr(err)
	}()

	state.SetStateDefaults()
	select {}
}
