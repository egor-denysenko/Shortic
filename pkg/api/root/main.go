package api_root

import (
	"net/http"
	apiUtility "shortic/pkg/api/utility"
	"shortic/pkg/dbservice"
	"strings"
	"unicode/utf8"
)

const ErrMarshallingJson = "Error provviding the shorted url."
const ErrWrongApiVersion = "Wrong Api Version Used."
const EndpointNotFound = "The provvided endpoint was not found"
const ShortedUrlNotValid = "The provvided shorted url is not valid."

func RedirectFromShortUrl(dbservice dbservice.ConsumerService) http.HandlerFunc{
  return func(w http.ResponseWriter,r *http.Request){

    // Check if only the supported method is called with the proper helper function
    if !apiUtility.ApiAllowedMethods(r.Method,apiUtility.AllowedMethods{"GET":{}}){
      w.WriteHeader(http.StatusNotImplemented)
    }

    // Confront if the Accept-Version header is compatible with the enpoint version
    if !apiUtility.ApiVersionHeaderCheck(r,"0.0.1"){
      http.Error(w, ErrWrongApiVersion, http.StatusGone)
    }

    urlPathSlice := strings.Split(r.URL.Path, "/")
    // Check If the url is containing only the shorten id and nothing else
    if len(urlPathSlice) < 2 || len(urlPathSlice) > 3 {
      http.Error(w,EndpointNotFound,http.StatusNotFound)
      return
    }
    
    if utf8.RuneCountInString(urlPathSlice[1]) < 7 || utf8.RuneCountInString(urlPathSlice[1]) > 7{
      http.Error(w,ShortedUrlNotValid,http.StatusBadRequest)
      return
    }
  } 
}

