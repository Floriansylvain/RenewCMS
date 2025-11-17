package server

import (
	"RenewCMS/api"
	"RenewCMS/main/route"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

var possibleEnvFileLocations = []string{".env", "../.env"}
var envVarsToLoad = []string{
	"HOST",
	"PORT",
	"JWT_SECRET",
	"ENVIRONMENT",
	"CORS_ALLOWED_ORIGINS",
	"DB_FILE",
	"SMTP_EMAIL",
	"SMTP_PASSWORD",
	"SMTP_HOST",
	"SMTP_PORT",
}

func initEnvVariables() {
	var err error
	for _, envLocation := range possibleEnvFileLocations {
		err = godotenv.Load(envLocation)
		if err == nil {
			break
		}
	}
	if err != nil {
		fmt.Println("WARNING: Could not load any .env file")
	}

	for _, envVar := range envVarsToLoad {
		if _, ok := os.LookupEnv(envVar); !ok {
			panic(fmt.Sprintf("Environment variable %s is not set", envVar))
		}
	}
}

func InitServer() *chi.Mux {
	initEnvVariables()
	api.InitContainer()
	api.InitValidator()
	route.InitJwt()
	return route.InitRoutes()
}

func StartServer(router *chi.Mux) error {
	fmt.Println("Server starting on :" + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
	return err
}
