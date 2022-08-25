package api

import (
	"net/http"
	"os"
	url_shortener "shortic/pkg/api/data/shortener"
	health_root "shortic/pkg/api/health"

	api_root "shortic/pkg/api/root"
	_ "shortic/docs"
	"shortic/pkg/dbservice"
  "log"
	//httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)


func ServeRestApi(dbservice *dbservice.DatabaseService) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", api_root.RedirectFromShortUrl(dbservice))
	mux.HandleFunc("/health/", health_root.RootHealth)

  // not implemented 
	mux.HandleFunc("/health/redis", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/data/shortener", url_shortener.FormUrlShortenerHandler(dbservice))

  /*mux.HandleFunc("/docs/swagger/",httpSwagger.Handler(
    httpSwagger.URL("doc.json"), //The url pointing to API definition
  ))*/

	log.Println("Starting to serve rest api on port",os.Getenv("API_PORT"))
  return http.ListenAndServe(":" + os.Getenv("API_PORT"), mux)
}
