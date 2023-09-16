package routers

import (
	"github.com/montvlein/DH_Backend4_final/cmd/server/handler"
	"github.com/montvlein/DH_Backend4_final/cmd/server/middleware"

	"github.com/gin-gonic/gin"
)

func SetupPatientsRouter(router *gin.Engine, patientsHandler *handler.PatientsHandler, authMidd *middleware.Auth) {
	patientsGroup := router.Group("/patients")

	patientsGroup.POST("/", authMidd.AuthHeader, patientsHandler.CreatePatient)
	patientsGroup.GET("/:id", patientsHandler.GetPatientByID)
	patientsGroup.GET("/list", patientsHandler.GetPatientList)
	patientsGroup.PUT("/:id", authMidd.AuthHeader, patientsHandler.PutPatient)
	patientsGroup.PATCH("/:id", authMidd.AuthHeader, patientsHandler.PatchPatient)
	patientsGroup.DELETE("/:id", authMidd.AuthHeader, patientsHandler.DeletePatient)
}
