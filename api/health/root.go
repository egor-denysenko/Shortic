package health_root

import (
	"fmt"
	"net/http"
	apiUtility "shortic/api/utility"
)

func RootHealth (w http.ResponseWriter, r *http.Request) {
  if !apiUtility.ApiVersionHeaderCheck(r,"0.0.1"){
    fmt.Fprintln(w,"Wrong api version")
    return
  }
  w.Write([]byte("Welcome from Shortic!"))
}
