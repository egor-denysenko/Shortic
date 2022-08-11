package api

import (
	"net/http"
	url_shortener "shortic/pkg/api/data/shortener"
	health_root "shortic/pkg/api/health"
	"shortic/pkg/dbservice"
  api_root "shortic/pkg/api/root"
)

func ServeRestApi(dbservice *dbservice.DatabaseService) error {
	mux := http.NewServeMux()
  mux.HandleFunc("/",api_root.RedirectFromShortUrl(dbservice))
	mux.HandleFunc("/health/", health_root.RootHealth)
	mux.HandleFunc("/health/redis", func(w http.ResponseWriter, r *http.Request) {})
  mux.HandleFunc("/data/shortener",url_shortener.FormUrlShortenerHandler(dbservice))
  //mux.HandleFunc()

	return http.ListenAndServe(":7878", mux)
}
