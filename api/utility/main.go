package utility

import (
	"net/http"
) 

//Utility function that is responsible in order to verify if the correct version of the api is called
func ApiVersionHeaderCheck(req *http.Request,endpointVersion string) bool{
  requestedApiVersion := req.Header.Get("Accept-Version")

  if requestedApiVersion != endpointVersion{
    return false
  }

  return true 
}
