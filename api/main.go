package api

import (
  "net/http"
  health_root "shortic/api/health"
)

func ServeRestApi() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/health/", health_root.RootHealth)
	mux.HandleFunc("/health/redis", func(w http.ResponseWriter, r *http.Request) {})

	return http.ListenAndServe(":7878", mux)
}
