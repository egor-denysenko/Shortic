package shortener

import (
	"net/http"
	apiUtility "shortic/api/utility"
)

func FormUrlShortener(w http.ResponseWriter, r *http.Request){
  // Check if only the supported method is called with the proper helper function
  if !apiUtility.ApiAllowedMethods(r.Method,apiUtility.AllowedMethods{"POST":{}}){
    w.WriteHeader(http.StatusNotImplemented)
  }

  if !apiUtility.ApiVersionHeaderCheck(r,"0.0.1"){
    w.WriteHeader(http.StatusGone)
    return
  }
}
