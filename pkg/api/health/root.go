package health_root

import (
	"net/http"
	apiUtility "shortic/pkg/api/utility"
)

// Shortic Health ... Retrieve if api is alive
// @Summary Check API Health
// @Description Chech API Health
// @Tags Health
// @Success 200
// @Failure 410 {object} object
// @Router /health/ [get]
func RootHealth(w http.ResponseWriter, r *http.Request) {
	if !apiUtility.ApiVersionHeaderCheck(r, "0.0.1") {
		w.WriteHeader(http.StatusGone)
		return
	}
	w.Write([]byte("Welcome from Shortic!"))
}
