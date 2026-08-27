package dashboard

import (
	"net/http"

	"github.com/Wlczak/vlastas-vpn/templates"
	"github.com/gin-gonic/gin"
)

type AdminDashValues struct {
	Locations []*ServerGeoLocation
}
type ServerGeoLocation struct {
	Code string
	Name string
}

func HandleAdminDashRoot(ctx *gin.Context) {
	adminDashValues := &AdminDashValues{
		Locations: []*ServerGeoLocation{{"location1", "Location1"}, {"location2", "Location2"}, {"location3", "Location3"}},
	}

	if ctx.Request.Method == http.MethodGet {
		const adminDashTemplate = "admin.tmpl"
		tmpl := templates.ParseTemplate(adminDashTemplate)
		tmpl.ExecuteTemplate(ctx.Writer, adminDashTemplate, adminDashValues)
	} else {
		location := ctx.PostForm("location")
		ctx.String(http.StatusOK, location)
	}

}
