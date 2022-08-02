package main

import (
	//"shortic/pkg/dbaccess"
	"log"
	restApi "shortic/api"
)

func main(){
  err := restApi.ServeRestApi()
  if err != nil{
    log.Default().Fatalln()
  }
}
