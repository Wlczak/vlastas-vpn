// @title           Vlasta's VPN api
// @version         1.0
// @description     API documentation for Vlasta's VPN.
// @host            vpn.vlastas.cc
// @BasePath        /api/v1
package api

import (
	"github.com/gin-gonic/gin"
)

func SetupApiRouter(apiRouter *gin.RouterGroup) *gin.RouterGroup {

	apiRouter.POST("/setLocation", HandleSetLocation)

	apiRouter.GET("/getCurrentLocation", HandleGetCurrentLocation)
	apiRouter.GET("/getLocationList", HandleGetLocationList)

	return apiRouter
}
