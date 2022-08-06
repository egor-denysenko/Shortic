package main

import (
	//"shortic/pkg/dbaccess"
	"log"
	restApi "shortic/api"
)

func main(){
  log.Println("Starting to serve rest api on port 7878")
  err := restApi.ServeRestApi()
  if err != nil{
    log.Default().Fatalln()
  }
}
