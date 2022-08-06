package shortener

import (
	"net/http"
	apiUtility "shortic/api/utility"
	"shortic/pkg/dbservice"
)

//Wrapper in order to accept and use the redis connection provvided.
func FormUrlShortenerHandler(dbservice *dbservice.QueueService) http.HandlerFunc{
   return func (w http.ResponseWriter, r *http.Request){
    // Check if only the supported method is called with the proper helper function
    if !apiUtility.ApiAllowedMethods(r.Method,apiUtility.AllowedMethods{"POST":{}}){
      w.WriteHeader(http.StatusNotImplemented)
    }
    // Confront if the Accept-Version header is compatible with the enpoint version
    if !apiUtility.ApiVersionHeaderCheck(r,"0.0.1"){
      w.WriteHeader(http.StatusGone)
    }

     

  }
}
