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

type ErrorResponse struct {
	Msg string
}

func HandleSetLocation(ctx *gin.Context) {
	type SetLocationRequest struct {
		LocationCode string
	}

	locationRequestByte, err := io.ReadAll(ctx.Request.Body)
	h.ChckErr(err)

	locationList := state.GetServerState().LocationList
	locationCodeList := make([]string, len(locationList))
	for i, location := range locationList {
		locationCodeList[i] = location.Code
	}

	setLocationRequest := &SetLocationRequest{}
	err = json.Unmarshal(locationRequestByte, &setLocationRequest)
	h.ChckErr(err)

	if slices.Contains(locationCodeList, setLocationRequest.LocationCode) {
		state.SetServerLocationByCode(setLocationRequest.LocationCode)
		ctx.JSON(http.StatusOK, setLocationRequest)
	} else {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{Msg: "Could not find " + setLocationRequest.LocationCode + " given location in the server list"})
	}

}

func HandleGetLocation(ctx *gin.Context) {
	type GetLocationResponse struct {
		Location state.ServerGeoLocation
	}
	location := state.GetServerState().CurrentLocation
	getLocationResponse := &GetLocationResponse{
		Location: location,
	}
	ctx.JSON(http.StatusOK, getLocationResponse)
}

func HandleGetLocationList(ctx *gin.Context) {
	s := state.GetServerState()
	ctx.JSON(http.StatusOK, s.LocationList)
}
