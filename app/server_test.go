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
	league   League
}

func (s *StubPlayerStore) GetLeague() (result League) {
	return s.league
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	server := NewPlayerServer(&StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	})

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
	var server = NewPlayerServer(store)

	t.Run("returns accepted on POST", func(t *testing.T) {
		var request = newPostWinRequest("Pepper")
		var respone = httptest.NewRecorder()

		server.ServeHTTP(respone, request)

		assertStatusCode(t, respone, http.StatusAccepted)
		assertWinCalls(t, store.winCalls, []string{"Pepper"})
	})
}

func TestLeague(t *testing.T) {
	want := League{
		{Name: "Cleo", Wins: 32},
		{Name: "Chris", Wins: 20},
		{Name: "Tiest", Wins: 14},
	}
	var server = NewPlayerServer(&StubPlayerStore{league: want})
	var response = httptest.NewRecorder()
	var request = newGetLeagueRequest()

	server.ServeHTTP(response, request)

	got, err := getLeagueFromResponse(response)

	assertNoError(t, err)
	assertPlayers(t, got, want)
	assertStatusCode(t, response, http.StatusOK)
	assertContentIsJson(t, response)
}

func getLeagueFromResponse(response *httptest.ResponseRecorder) (league League, err error) {
	league, err = NewLeague(response.Body)
	return
}

func assertScore(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("score got %d, want %d", got, want)
	}
}

func assertContentIsJson(t testing.TB, response *httptest.ResponseRecorder) {
	got := response.Header().Get("content-type")
	want := "application/json"

	if got != want {
		t.Fatalf("content type got %q, want %q", got, want)
	}
}

func assertPlayers(t testing.TB, got League, want League) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("league got %v want %v", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("did not expect an error: %#v", err)
	}
}

func newGetLeagueRequest() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/league", nil)
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
