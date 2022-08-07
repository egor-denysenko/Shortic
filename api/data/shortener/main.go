package shortener

import (
  "encoding/json"
	"net/http"
	apiUtility "shortic/api/utility"
	"shortic/pkg/dbservice"
  "net/url"
)

type postUrlBody struct{
  url string
}

//Wrapper in order to accept and use the redis connection provvided.
func FormUrlShortenerHandler(dbservice *dbservice.QueueService) http.HandlerFunc{
   return func (w http.ResponseWriter, r *http.Request){

    // Check if only the supported method is called with the proper helper function
    if !apiUtility.ApiAllowedMethods(r.Method,apiUtility.AllowedMethods{"POST":{}}){
      w.WriteHeader(http.StatusNotImplemented)
    }

    // Confront if the Accept-Version header is compatible with the enpoint version
    if !apiUtility.ApiVersionHeaderCheck(r,"0.0.1"){
      http.Error(w, "Wrong API Version Used", http.StatusGone)
      //w.WriteHeader(http.StatusGone)
    }

    var urlBodyStruct postUrlBody

    // Decode body payload into golang struct
    errDecodingJson := json.NewDecoder(r.Body).Decode(&urlBodyStruct)
    if errDecodingJson != nil {
        http.Error(w, errDecodingJson.Error(), http.StatusBadRequest)
        return
    }

    if !isUrlValid(urlBodyStruct.url){
      w.WriteHeader(http.StatusInternalServerError) 
    }   



  }
}

// Helper funciton that checks if the provvided url is a valid url
func isUrlValid(fullUserUrl string) bool{
  _, err := url.ParseRequestURI(fullUserUrl) 

  if err != nil {
    return false
  }

  return true

}

func shortenUrl(fullUserUrl string) string{
  return ""
}
