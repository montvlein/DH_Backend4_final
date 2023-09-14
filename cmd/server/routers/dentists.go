package routers

import (
	"github.com/montvlein/DH_Backend4_final/cmd/server/handler"
	"github.com/montvlein/DH_Backend4_final/cmd/server/middleware"

	"github.com/gin-gonic/gin"
)

func SetupDentistsRouter(router *gin.Engine, dentistsHandler *handler.DentistsHandler, authMidd *middleware.Auth) {
	dentistsGroup := router.Group("/dentists")

	dentistsGroup.POST("/", authMidd.AuthHeader, dentistsHandler.CreateDentist)
	dentistsGroup.GET("/:id", dentistsHandler.GetDentistByID)
	dentistsGroup.PUT("/:id", authMidd.AuthHeader, dentistsHandler.PutDentist)
	dentistsGroup.PATCH("/:id", authMidd.AuthHeader, dentistsHandler.PatchDentist)
	dentistsGroup.DELETE("/:id", authMidd.AuthHeader, dentistsHandler.DeleteDentist)
}
