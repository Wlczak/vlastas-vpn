package api

import (
	"github.com/Wlczak/vlastas-vpn/api/dashboard"
	"github.com/gin-gonic/gin"
)

func RunApi() {
	r := gin.Default()

	r.Any("/admin", dashboard.HandleAdminDashRoot)
	r.Run()
}
