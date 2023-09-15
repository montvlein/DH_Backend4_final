package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/montvlein/DH_Backend4_final/cmd/server/handler"
	"github.com/montvlein/DH_Backend4_final/cmd/server/middleware"
)

func SetupAppointmentsRouter(router *gin.Engine, appointmentsHandler *handler.AppointmentsHandler, authMiddleware *middleware.Auth) {
	appointmentsGroup := router.Group("/appointments")

	appointmentsGroup.POST("/", authMiddleware.AuthHeader, appointmentsHandler.CreateAppointment)
	appointmentsGroup.GET("/:id", appointmentsHandler.GetAppointmentByID)
	appointmentsGroup.GET("/:id/details", appointmentsHandler.GetAppointmentByIDWithEntities)
	appointmentsGroup.PUT("/:id", authMiddleware.AuthHeader, appointmentsHandler.PutAppointment)
	appointmentsGroup.PATCH("/:id", authMiddleware.AuthHeader, appointmentsHandler.PatchAppointment)
	appointmentsGroup.DELETE("/:id", authMiddleware.AuthHeader, appointmentsHandler.DeleteAppointment)
	appointmentsGroup.POST("/create-by-dni-and-license", authMiddleware.AuthHeader, appointmentsHandler.CreateAppointmentByDNIAndLicense)
	appointmentsGroup.GET("/patient", appointmentsHandler.GetAppointmentsByPatientDNI)
}
