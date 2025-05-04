package main

import (
	"EffectiveMobileTest/internal/controllers"
	"EffectiveMobileTest/internal/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	log.SetReportCaller(true)
	log.Debug("logger debug mode enabled")

	dbErr := database.SetupDatabase()
	if dbErr != nil {
		log.Error(dbErr)
	}
	r := mux.NewRouter()

	Routes(r)

	log.Fatal(http.ListenAndServe(":10000", r))
}

func Routes(r *mux.Router) {

	r.StrictSlash(true)
	r.HandleFunc("/data", controllers.DataCreate).Methods("POST")
	r.HandleFunc("/data", controllers.DataGet).Methods("GET")
	r.HandleFunc("/data", controllers.DataDelete).Methods("DELETE")
	r.HandleFunc("/data", controllers.DataUpdate).Methods("PUT")
}
