package dashboard

import (
	"net/http"

	h "github.com/Wlczak/vlastas-vpn/helpers"
	"github.com/Wlczak/vlastas-vpn/state"
	"github.com/Wlczak/vlastas-vpn/templates"
	"github.com/gin-gonic/gin"
)

type AdminDashValues struct {
	Locations       []state.ServerGeoLocation
	CurrentLocation state.ServerGeoLocation
}

func HandleAdminDashRoot(ctx *gin.Context) {
	if ctx.Request.Method == http.MethodPost {
		locationCode := ctx.PostForm("location")
		state.SetServerLocationByCode(locationCode)
		// ctx.String(http.StatusOK, location)
		// ctx.Redirect(http.StatusTemporaryRedirect, "/admin")
	}

	adminDashValues := &AdminDashValues{
		Locations:       state.GetServerState().LocationList,
		CurrentLocation: state.GetServerState().CurrentLocation,
	}

	const adminDashTemplate = "admin.tmpl"
	tmpl := templates.ParseTemplate(adminDashTemplate)
	err := tmpl.ExecuteTemplate(ctx.Writer, adminDashTemplate, adminDashValues)
	h.ChckErr(err)
}
