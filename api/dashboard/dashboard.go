package dashboard

import (
	"net/http"

	"github.com/Wlczak/vlastas-vpn/state"
	"github.com/Wlczak/vlastas-vpn/templates"
	"github.com/gin-gonic/gin"
)

type AdminDashValues struct {
	Locations       []*state.ServerGeoLocation
	CurrentLocation string
}

func HandleAdminDashRoot(ctx *gin.Context) {
	if ctx.Request.Method == http.MethodPost {
		location := ctx.PostForm("location")
		state.SetServerLocation(state.ServerGeoLocation{Code: location, Name: location})
		// ctx.String(http.StatusOK, location)
		// ctx.Redirect(http.StatusTemporaryRedirect, "/admin")
	}

	adminDashValues := &AdminDashValues{
		Locations:       []*state.ServerGeoLocation{{"location1", "Location1"}, {"location2", "Location2"}, {"location3", "Location3"}},
		CurrentLocation: state.GetServerState().CurrentLocation.Name,
	}

	const adminDashTemplate = "admin.tmpl"
	tmpl := templates.ParseTemplate(adminDashTemplate)
	tmpl.ExecuteTemplate(ctx.Writer, adminDashTemplate, adminDashValues)

}
