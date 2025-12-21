package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	player := "Pepper"

	wins := 3
	for range wins {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	t.Run("get score", func(t *testing.T) {

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assertStatusCode(t, response, http.StatusOK)
		assertResponseBody(t, response, strconv.Itoa(wins))
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()

		server.ServeHTTP(response, newGetLeagueRequest())

		assertStatusCode(t, response, http.StatusOK)

		want := []Player{{Name: player, Wins: wins}}
		got, _ := getLeagueFromResponse(response)
		assertPlayers(t, got, want)
	})

}
