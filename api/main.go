package api

import (
	"net/http"
	url_shortener "shortic/api/data/shortener"
	health_root "shortic/api/health"
)

func ServeRestApi() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/health/", health_root.RootHealth)
	mux.HandleFunc("/health/redis", func(w http.ResponseWriter, r *http.Request) {})
  mux.HandleFunc("/data/shortener",url_shortener.FormUrlShortener)
  //mux.HandleFunc()

	return http.ListenAndServe(":7878", mux)
}
