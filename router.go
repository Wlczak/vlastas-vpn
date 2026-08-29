package main

import (
	"github.com/Wlczak/vlastas-vpn/api"
	"github.com/Wlczak/vlastas-vpn/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	controllers.SetupControllerRouter(r.Group("/"))
	api.SetupApiRouter(r.Group("/api/v1"))

	return r
}
