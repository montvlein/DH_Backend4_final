package main

import (
	"net/http"
	"os"

	"github.com/montvlein/DH_Backend4_final/cmd/server/config"
	"github.com/montvlein/DH_Backend4_final/cmd/server/external/database"
	"github.com/montvlein/DH_Backend4_final/cmd/server/handler"
	"github.com/montvlein/DH_Backend4_final/cmd/server/middleware"
	"github.com/montvlein/DH_Backend4_final/cmd/server/routers"
	"github.com/montvlein/DH_Backend4_final/internal/dentists"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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

	routers.SetupDentistsRouter(router, dentistsHandler, authMidd)

	err = router.Run()

	if err != nil {
		panic(err)
	}
}
