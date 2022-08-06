package api

import (
	"net/http"
	url_shortener "shortic/api/data/shortener"
	health_root "shortic/api/health"
	"shortic/pkg/dbservice"
)

func ServeRestApi(dbservice *dbservice.QueueService) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/health/", health_root.RootHealth)
	mux.HandleFunc("/health/redis", func(w http.ResponseWriter, r *http.Request) {})
  mux.HandleFunc("/data/shortener",url_shortener.FormUrlShortenerHandler(dbservice))
  //mux.HandleFunc()

	return http.ListenAndServe(":7878", mux)
}
