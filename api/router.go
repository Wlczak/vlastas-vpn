package api

import (
	"github.com/Wlczak/vlastas-vpn/api/dashboard"
	"github.com/gin-gonic/gin"
)

func RunApi() {
	r := gin.Default()

	r.Any("/admin", dashboard.HandleAdminDashRoot)

	apiRouter := r.Group("/api")

	{
		apiRouter.POST("/setLocation", HandleSetLocation)
		apiRouter.GET("/getLocation", HandleGetLocation)
	}

	r.Run()
}
