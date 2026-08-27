// @title           Vlasta's VPN api
// @version         1.0
// @description     API documentation for Vlasta's VPN.
// @host            vpn.vlastas.cc
// @BasePath        /api/v1
package api

import (
	"github.com/Wlczak/vlastas-vpn/api/dashboard"
	"github.com/gin-gonic/gin"
)

func RunApi() {
	r := gin.Default()

	r.Any("/admin", dashboard.HandleAdminDashRoot)

	apiRouter := r.Group("/api/v1")

	{
		apiRouter.POST("/setLocation", HandleSetLocation)

		apiRouter.GET("/getLocation", HandleGetLocation)
		apiRouter.GET("/getLocationList", HandleGetLocationList)
	}

	r.Run()
}
