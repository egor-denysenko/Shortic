package health_root

import (
  "net/http"
)

func RootHealth (w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Welcome from Shortic!"))
}
