package main

import (
	"RenewCMS/api"
	validator "RenewCMS/api/pkg/validator"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

var possibleEnvFileLocations = []string{".env", "../.env"}
var envVarsToLoad = []string{
	"HOST",
	"PORT",
	"JWT_SECRET",
	"ENVIRONMENT",
	"CORS_ALLOWED_ORIGINS",
	"DB_FILE",
	"UPLOAD_DIR",
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

func initServer() *chi.Mux {
	initEnvVariables()

	container := api.InitContainer()
	validator.InitValidator()

	return api.InitRoutes(container)
}

func startServer(router *chi.Mux) error {
	fmt.Println("Server starting on http://" + os.Getenv("HOST"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
	return err
}

func main() {
	router := initServer()
	err := startServer(router)
	if err != nil {
		panic(err)
	}
}
