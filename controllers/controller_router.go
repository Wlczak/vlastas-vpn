package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetupControllerRouter(controllerRouter *gin.RouterGroup) *gin.RouterGroup {

	// controllerRouter.Any("/admin", dashboard.HandleAdminDashRoot)

	return controllerRouter
}
