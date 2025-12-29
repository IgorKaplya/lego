package app_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/IgorKaplya/lego/app"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanupDatabase := createTempDatabase(t)
	defer cleanupDatabase()

	store, err := app.NewFileSystemPlayerStore(database)
	assertNoError(t, err)

	game := app.NewGame(app.BlindAlerterFun(app.Alerter), store)
	server := mustMakePlayerServer(t, store, game)
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

		want := app.League{{Name: player, Wins: wins}}
		got, _ := getLeagueFromResponse(response)
		assertPlayers(t, got, want)
	})

}

func createTempDatabase(t testing.TB) (app.FileDatabase, func()) {
	file, errCreate := os.CreateTemp("", "db")
	assertNoError(t, errCreate)

	file.Write([]byte("[]"))

	return file, func() {
		assertNoError(t, file.Close())
		assertNoError(t, os.Remove(file.Name()))
	}
}
