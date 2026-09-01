package api

import (
	"encoding/json"
	"io"
	"net/http"
	"slices"

	h "github.com/Wlczak/vlastas-vpn/helpers"
	"github.com/Wlczak/vlastas-vpn/state"
	"github.com/gin-gonic/gin"
)

func HandleSetLocation(ctx *gin.Context) {
	type SetLocationRequest struct {
		LocationHostname string
	}

	locationRequestByte, err := io.ReadAll(ctx.Request.Body)
	h.PanicChckErr(err)

	locationList := state.GetServerState().LocationList
	locationCodeList := make([]string, len(locationList))
	for i, location := range locationList {
		locationCodeList[i] = location.Hostname
	}

	setLocationRequest := &SetLocationRequest{}
	err = json.Unmarshal(locationRequestByte, &setLocationRequest)
	h.PanicChckErr(err)

	if slices.Contains(locationCodeList, setLocationRequest.LocationHostname) {
		state.SetServerLocationByHostname(setLocationRequest.LocationHostname)
		ctx.JSON(http.StatusOK, setLocationRequest)
	} else {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{Msg: "Could not find " + setLocationRequest.LocationHostname + " given location in the server list"})
	}

}

// Get current location godoc
// @Summary      Get current server location
// @Description  Retrieves the currently set server location from state
// @Tags         location
// @Produce      json
// @Success      200  {object}  state.ServerGeoLocation
// @Failure      404  {object}  ErrorResponse
// @Router       /getLocation [get]
func HandleGetLocation(ctx *gin.Context) {
	type GetLocationResponse struct {
		Location state.MullvadServerLocation
	}
	location := state.GetServerState().CurrentLocation
	getLocationResponse := &GetLocationResponse{
		Location: location,
	}
	ctx.JSON(http.StatusOK, getLocationResponse)
}

func HandleGetLocationList(ctx *gin.Context) {
	s := state.GetServerState()
	freshMullvadServerLocationList := fetchMullvadLocationList()

	s.LocationList = freshMullvadServerLocationList
	ctx.JSON(http.StatusOK, s.LocationList)
}

func fetchMullvadLocationList() []state.MullvadServerLocation {
	url := "https://api.mullvad.net/www/relays/all" // TODO: Move to constants envetually
	resp, err := http.Get(url)
	h.PanicChckErr(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	h.PanicChckErr(err)

	mullvadServerLocationList := []state.MullvadServerLocation{}
	err = json.Unmarshal(body, &mullvadServerLocationList)
	h.PanicChckErr(err)

	return mullvadServerLocationList
}
