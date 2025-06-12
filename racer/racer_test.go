package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("when URLs respond faster than 10s", func(t *testing.T) {
		var slowServer = CreateTestServer(20 * time.Millisecond)
		var fastServer = CreateTestServer(time.Nanosecond)

		defer slowServer.Close()
		defer fastServer.Close()

		var got, err = Racer(slowServer.URL, fastServer.URL)

		if err != nil {
			t.Fatalf("Unexpected error %#v", err)
		}

		if got != fastServer.URL {
			t.Errorf("got %q want %q", got, fastServer.URL)
		}
	})
	t.Run("when URLs respond slower than 10s", func(t *testing.T) {
		var slowServer = CreateTestServer(25 * time.Millisecond)
		defer slowServer.Close()

		var _, err = ConfigurableRacer(slowServer.URL, slowServer.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but got nil")
		}
	})
}

func CreateTestServer(responseDelay time.Duration) (result *httptest.Server) {
	result = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(responseDelay)
		w.WriteHeader(http.StatusOK)
	}))
	return result
}
