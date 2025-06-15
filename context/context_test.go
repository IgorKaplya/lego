package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StoreSpy struct {
	data        string
	dataTimeout time.Duration
	cancelled   bool
}

func (s *StoreSpy) Cancel() {
	s.cancelled = true
}

func (s *StoreSpy) Fetch(context context.Context) string {
	time.Sleep(s.dataTimeout)

	select {
	case <-context.Done():
		s.Cancel()
		return ""
	default:
		return s.data
	}
}

func TestServer(t *testing.T) {
	t.Run("returns data", func(t *testing.T) {
		var want = "hello, world"
		var store = &StoreSpy{data: want, dataTimeout: 0, cancelled: false}
		var server http.HandlerFunc = Server(store)

		var request = httptest.NewRequest(http.MethodGet, "/", nil)
		var response = httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got string = response.Body.String()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("store cancels on request cancel", func(t *testing.T) {
		var want = "hello, world"
		var store = &StoreSpy{data: want, dataTimeout: 25 * time.Millisecond, cancelled: false}
		var server http.HandlerFunc = Server(store)

		var request = httptest.NewRequest(http.MethodGet, "/", nil)
		var cancellingContext, cancel = context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingContext)

		var response = httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if !store.cancelled {
			t.Error("Store was not cancelled")
		}

		if response.Body.String() != "" {
			t.Errorf("got %q want %q", response.Body.String(), "")
		}
	})
}
