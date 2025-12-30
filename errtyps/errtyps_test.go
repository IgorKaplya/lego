package errtyps_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IgorKaplya/lego/errtyps"
)

func TestDumbGetter(t *testing.T) {
	t.Run("returns status error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusTeapot)
		}))
		defer server.Close()

		url := server.URL
		_, err := errtyps.DumbGetter(url)

		if err == nil {
			t.Fatal("expected err")
		}

		got, isStatusError := err.(errtyps.BadStatusError)
		if !isStatusError {
			t.Fatalf("expected BadStatusError, %T", err)
		}
		want := errtyps.BadStatusError{URL: url, Status: http.StatusTeapot}

		if got != want {
			t.Errorf("error got %q, want %q", got, want)
		}
	})
}
