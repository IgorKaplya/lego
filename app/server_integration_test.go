package app

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanupDatabase := createTempDatabase(t)
	defer cleanupDatabase()

	store, err := NewFileSystemPlayerStore(database)
	AssertNoError(t, err)

	server := NewPlayerServer(store)
	player := "Pepper"

	wins := 3
	for range wins {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	t.Run("get score", func(t *testing.T) {

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		AssertStatusCode(t, response, http.StatusOK)
		AssertResponseBody(t, response, strconv.Itoa(wins))
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()

		server.ServeHTTP(response, newGetLeagueRequest())

		AssertStatusCode(t, response, http.StatusOK)

		want := League{{Name: player, Wins: wins}}
		got, _ := GetLeagueFromResponse(response)
		AssertPlayers(t, got, want)
	})

}

func createTempDatabase(t testing.TB) (FileDatabase, func()) {
	file, errCreate := os.CreateTemp("", "db")
	AssertNoError(t, errCreate)

	file.Write([]byte("[]"))

	return file, func() {
		AssertNoError(t, file.Close())
		AssertNoError(t, os.Remove(file.Name()))
	}
}
