package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	server := &PlayerServer{store: &StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response, http.StatusOK)
		assertResponseBody(t, response, "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response, http.StatusOK)
		assertResponseBody(t, response, "10")
	})

	t.Run("returns 404 for missing players", func(t *testing.T) {
		request := newGetScoreRequest("Appolo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	var store = &StubPlayerStore{
		scores:   map[string]int{},
		winCalls: []string{},
	}
	var server = &PlayerServer{store: store}

	t.Run("returns accepted on POST", func(t *testing.T) {
		var request = newPostWinRequest("Pepper")
		var respone = httptest.NewRecorder()

		server.ServeHTTP(respone, request)

		assertStatusCode(t, respone, http.StatusAccepted)
		assertWinCalls(t, store.winCalls, []string{"Pepper"})
	})
}

func newPostWinRequest(name string) *http.Request {
	var result, _ = http.NewRequest(http.MethodPost, "/players/"+name, nil)
	return result
}

func assertWinCalls(t testing.TB, got []string, want []string) {
	t.Helper()

	wantLen := len(want)
	gotLen := len(got)

	if wantLen != gotLen {
		t.Errorf("want wins len %d, got %d", wantLen, gotLen)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want wins %v, got %v", want, got)
	}
}

func assertStatusCode(t *testing.T, response *httptest.ResponseRecorder, want int) {
	t.Helper()

	got := response.Code

	if got != want {
		t.Errorf("got status %d want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, response *httptest.ResponseRecorder, want string) {
	got := response.Body.String()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func newGetScoreRequest(name string) *http.Request {
	return httptest.NewRequest(http.MethodGet, "/players/"+name, nil)
}
