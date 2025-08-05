package routes

import (
	"app/app/modules"

	"github.com/gin-gonic/gin"
)

func Staff(router *gin.RouterGroup) {
	module := modules.New()
	staff := router.Group("")
	{
		staff.POST("/create", module.Staff.Ctl.Create)
		staff.POST("/login", module.Staff.Ctl.Login)
	}
}
