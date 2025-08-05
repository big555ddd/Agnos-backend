package routes

import (
	"app/app/middleware"
	"app/app/modules"

	"github.com/gin-gonic/gin"
)

func Patient(router *gin.RouterGroup) {
	module := modules.New()
	amd := middleware.AuthMiddleware()
	patient := router.Group("")
	{
		patient.GET("/search/:id", module.Patient.Ctl.GetPatient)
		patient.GET("/search", amd, module.Patient.Ctl.List)
	}
}
