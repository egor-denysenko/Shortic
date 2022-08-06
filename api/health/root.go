package health_root

import (
	"net/http"
	apiUtility "shortic/api/utility"
)

func RootHealth (w http.ResponseWriter, r *http.Request) {
  if !apiUtility.ApiVersionHeaderCheck(r,"0.0.1"){
    w.WriteHeader(http.StatusGone)
    return
  }
  w.Write([]byte("Welcome from Shortic!"))
}
