package main

import (
	"flag"
	"log"
	restApi "shortic/pkg/api"
	"shortic/pkg/dbservice"
	"shortic/pkg/deployment"
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
  
	dbservice := dbservice.QueueServiceFactory()
	dberr := dbservice.Connect()

	if dberr != nil {
		log.Default().Fatalln("Failed Database Connection")
	}

	log.Println("Starting to serve rest api on port 7878")

	err := restApi.ServeRestApi(dbservice)

	if err != nil {
		log.Default().Fatalln("Failed Exposing REST API")
	}
}
