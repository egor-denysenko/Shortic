package health_root

import (
  "testing"
  "net/http"
  "net/http/httptest"
)

func TestRootHealth(t *testing.T) {
  t.Run("Verify the correct message is displayed",func(t *testing.T) {

    request, _ := http.NewRequest(http.MethodGet, "/health/", nil)
		response := httptest.NewRecorder()

		RootHealth(response, request)

		got := response.Body.String()
		want := "Welcome from Shortic!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
    }

  })
}
