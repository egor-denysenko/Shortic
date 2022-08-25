package main

import (
	"flag"
	"log"
	restApi "shortic/pkg/api"
	"shortic/pkg/dbservice"
	"shortic/pkg/deployment"

	"github.com/joho/godotenv"
)

// @title   Shortic API documentation
// @version 0.0.1

// @host     localhost:7878
// @BasePath /

func main() {
  flag.BoolVar(&deployment.SoftwareInfo, "h", false, "log specs about the software")
  flag.Parse()

  if deployment.SoftwareInfo {
    deployment.LogSoftwareInfo()
  }
  errEnv := godotenv.Load()
  if errEnv != nil {
    log.Fatal("Error loading .env file")
  }

	dbservice := dbservice.QueueServiceFactory()
	dberr := dbservice.Connect()

	if dberr != nil {
		log.Default().Fatalln("Failed Database Connection")
	}


	ApiErr := restApi.ServeRestApi(dbservice)

	if ApiErr != nil {
    log.Default().Fatalln("Failed Exposing REST API:",ApiErr)
	}
}
