package deployment

import (
	"log"
)

var version string
var commitHash string
var buildUser string
var buildTime string

var SoftwareInfo bool

func LogSoftwareInfo(){
  log.Print("Version:\t",version)
  log.Print("Commit hash:\t",commitHash)
  log.Print("Built by:\t",buildUser)
  log.Print("Built at:\t",buildTime)
}
