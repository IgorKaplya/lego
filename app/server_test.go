package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	server := NewPlayerServer(&StubPlayerStore{
		Scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	})

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatusCode(t, response, http.StatusOK)
		AssertResponseBody(t, response, "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatusCode(t, response, http.StatusOK)
		AssertResponseBody(t, response, "10")
	})

	t.Run("returns 404 for missing players", func(t *testing.T) {
		request := newGetScoreRequest("Appolo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatusCode(t, response, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	var store = &StubPlayerStore{
		Scores:   map[string]int{},
		WinCalls: []string{},
	}
	var server = NewPlayerServer(store)

	t.Run("returns accepted on POST", func(t *testing.T) {
		var request = newPostWinRequest("Pepper")
		var respone = httptest.NewRecorder()

		server.ServeHTTP(respone, request)

		AssertStatusCode(t, respone, http.StatusAccepted)
		AssertWinCalls(t, store.WinCalls, []string{"Pepper"})
	})
}

func TestLeague(t *testing.T) {
	want := League{
		{Name: "Cleo", Wins: 32},
		{Name: "Chris", Wins: 20},
		{Name: "Tiest", Wins: 14},
	}
	var server = NewPlayerServer(&StubPlayerStore{League: want})
	var response = httptest.NewRecorder()
	var request = newGetLeagueRequest()

	server.ServeHTTP(response, request)

	got, err := GetLeagueFromResponse(response)

	AssertNoError(t, err)
	AssertPlayers(t, got, want)
	AssertStatusCode(t, response, http.StatusOK)
	AssertContentIsJson(t, response)
}

func newGetLeagueRequest() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/league", nil)
}

func newPostWinRequest(name string) *http.Request {
	var result, _ = http.NewRequest(http.MethodPost, "/players/"+name, nil)
	return result
}

func newGetScoreRequest(name string) *http.Request {
	return httptest.NewRequest(http.MethodGet, "/players/"+name, nil)
}
