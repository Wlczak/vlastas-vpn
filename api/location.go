package api

import (
	"encoding/json"
	"io"
	"net/http"
	"slices"

	"github.com/Wlczak/vlastas-vpn/state"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Msg string
}

func HandleSetLocation(ctx *gin.Context) {
	type SetLocationRequest struct {
		Location string
	}

	locationByte, _ := io.ReadAll(ctx.Request.Body)
	setLocationRequest := &SetLocationRequest{}
	json.Unmarshal(locationByte, &setLocationRequest)

	if slices.Contains(state.GetServerState().LocationList, setLocationRequest.Location) {
		state.SetServerLocation(setLocationRequest.Location)
		ctx.JSON(http.StatusOK, setLocationRequest)
	} else {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{Msg: "Could not find given location in the server list"})
	}

}

func HandleGetLocation(ctx *gin.Context) {
	type GetLocationResponse struct {
		Location string
	}
	location := state.GetServerState().CurrentLocation
	getLocationResponse := &GetLocationResponse{
		Location: location,
	}
	ctx.JSON(http.StatusOK, getLocationResponse)
}
