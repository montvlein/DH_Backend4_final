package main

import (
	"net/http"
	"os"

	"github.com/montvlein/DH_Backend4_final/cmd/server/config"
	"github.com/montvlein/DH_Backend4_final/cmd/server/external/database"
	"github.com/montvlein/DH_Backend4_final/cmd/server/handler"
	"github.com/montvlein/DH_Backend4_final/cmd/server/middleware"
	"github.com/montvlein/DH_Backend4_final/cmd/server/routers"
	"github.com/montvlein/DH_Backend4_final/internal/appointments"
	"github.com/montvlein/DH_Backend4_final/internal/dentists"
	"github.com/montvlein/DH_Backend4_final/internal/patients"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/montvlein/DH_Backend4_final/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Grupo 8 Final Backend
// @version 1.0
// @description Solución al final de backend de Digital House - Grupo 8 por Felipe Monterrosa, Javier Triana, Fabricio Montivero, Gaston Diaz
// @termsOfService https://developers.ctd.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.ctd.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	if env == "local" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	cfg, err := config.NewConfig(env)

	if err != nil {
		panic(err)
	}

	authMidd := middleware.NewAuth(cfg.PublicConfig.PublicKey, cfg.PrivateConfig.SecretKey)

	router := gin.New()

	customRecovery := gin.CustomRecovery(middleware.RecoveryWithLog)

	router.Use(customRecovery)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"ok": "ok"})
	})

	postgresDatabase, err := database.NewPostgresSQLDatabase(cfg.PublicConfig.PostgresHost,
		cfg.PublicConfig.PostgresPort, cfg.PublicConfig.PostgresUser, cfg.PrivateConfig.PostgresPassword,
		cfg.PublicConfig.PostgresDatabase)

	if err != nil {
		panic(err)
	}

	myDatabase := database.NewDatabase(postgresDatabase)

	dentistsService := dentists.NewService(myDatabase)
	dentistsHandler := handler.NewDentistsHandler(dentistsService, dentistsService)

	patientsService := patients.NewService(myDatabase)
	patientsHandler := handler.NewPatientsHandler(patientsService, patientsService)

	appointmentsService := appointments.NewService(myDatabase)
	appointmentsHandler := handler.NewAppointmentsHandler(appointmentsService, appointmentsService)

	routers.SetupDentistsRouter(router, dentistsHandler, authMidd)
	routers.SetupPatientsRouter(router, patientsHandler, authMidd)
	routers.SetupAppointmentsRouter(router, appointmentsHandler, authMidd)

	err = router.Run()

	if err != nil {
		panic(err)
	}
}
