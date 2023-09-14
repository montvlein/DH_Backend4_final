package main

import (
	"DH_Backend4_final/cmd/server/config"
	"DH_Backend4_final/cmd/server/external/memory"
	"DH_Backend4_final/cmd/server/handler"
	"DH_Backend4_final/cmd/server/middleware"
	"DH_Backend4_final/internal/dentist"
	"DH_Backend4_final/internal/patient"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Certified Tech Developer
// @version 1.0
// @description This API Handle Products.
// @termsOfService https://developers.ctd.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.ctd.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	godotenv.Load()

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	cfg, err := config.NewConfig(env)

	if err != nil {
		panic(err)
	}

	authMidd := middleware.NewAuth(cfg.PublicConfig.PublicKey, cfg.PrivateConfig.SecretKey)

	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"ok": "ok"})
	})

	repositoryDentist, err := memory.NewServiceDentist("dentists.json")
	if err != nil {
		panic(err)
	}
	dentistsService := dentist.NewServiceDentist(repositoryDentist)
	dentistsHandler := handler.NewDentistsHandler(dentistsService, dentistsService)

	repositoryPatient, err := memory.NewServicePatient("patients.json")
	if err != nil {
		panic(err)
	}
	patientsService := patient.NewServicePatient(repositoryPatient)
	patientsHandler := handler.NewPatientsHandler(patientsService, patientsService)

	dentistsGroup := router.Group("/dentists")
	dentistsGroup.GET("/:id", dentistsHandler.GetDentistByID)
	dentistsGroup.PUT("/:id", authMidd.AuthHeader, dentistsHandler.PutDentist)

	patientsGroup := router.Group("/patients")
	patientsGroup.GET("/:id", patientsHandler.GetPatientByID)
	patientsGroup.PUT("/:id", authMidd.AuthHeader, patientsHandler.PutPatient)

	router.Run()

}
