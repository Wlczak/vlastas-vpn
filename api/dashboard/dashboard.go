package dashboard

import (
	"net/http"

	h "github.com/Wlczak/vlastas-vpn/helpers"
	"github.com/Wlczak/vlastas-vpn/state"
	"github.com/Wlczak/vlastas-vpn/templates"
	"github.com/gin-gonic/gin"
)

type AdminDashValues struct {
	Locations       []state.MullvadServerLocation
	CurrentLocation state.MullvadServerLocation
}

func HandleAdminDashRoot(ctx *gin.Context) {
	if ctx.Request.Method == http.MethodPost {
		locationHostname := ctx.PostForm("hostname")
		state.SetServerLocationByHostname(locationHostname)
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
