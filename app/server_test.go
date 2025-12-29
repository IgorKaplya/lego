package app_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/IgorKaplya/lego/app"
	"github.com/gorilla/websocket"
)

type stubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   app.League
}

func (s *stubPlayerStore) GetLeague() (result app.League) {
	return s.League
}

func (s *stubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

func (s *stubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

var dummyGameSpy = new(GameSpy)

func TestGETPlayers(t *testing.T) {
	server := mustMakePlayerServer(t, &stubPlayerStore{
		Scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}, dummyGameSpy)

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
	var store = &stubPlayerStore{
		Scores:   map[string]int{},
		WinCalls: []string{},
	}
	var server = mustMakePlayerServer(t, store, dummyGameSpy)

	t.Run("returns accepted on POST", func(t *testing.T) {
		var request = newPostWinRequest("Pepper")
		var respone = httptest.NewRecorder()

		server.ServeHTTP(respone, request)

		assertStatusCode(t, respone, http.StatusAccepted)
		assertWinCalls(t, store.WinCalls, []string{"Pepper"})
	})
}

func TestLeague(t *testing.T) {
	want := app.League{
		{Name: "Cleo", Wins: 32},
		{Name: "Chris", Wins: 20},
		{Name: "Tiest", Wins: 14},
	}
	var server = mustMakePlayerServer(t, &stubPlayerStore{League: want}, dummyGameSpy)
	var response = httptest.NewRecorder()
	var request = newGetLeagueRequest()

	server.ServeHTTP(response, request)

	got, err := getLeagueFromResponse(response)

	assertNoError(t, err)
	assertPlayers(t, got, want)
	assertStatusCode(t, response, http.StatusOK)
	assertContentIsJson(t, response)
}

func TestGame(t *testing.T) {
	t.Run("returns OK", func(t *testing.T) {
		dummyStore := &stubPlayerStore{}
		server := mustMakePlayerServer(t, dummyStore, dummyGameSpy)
		response := httptest.NewRecorder()
		request := newGetGameRequest()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response, http.StatusOK)
	})
}

func TestWs(t *testing.T) {
	t.Run("stores winner", func(t *testing.T) {
		store := &stubPlayerStore{}
		winner := "baba"
		wantAlert := "Blind is 100"
		game := &GameSpy{blindAlert: []byte(wantAlert)}
		server := httptest.NewServer(mustMakePlayerServer(t, store, game))
		defer server.Close()

		wsUrl := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

		ws := mustDialWs(t, wsUrl)
		defer ws.Close()

		assertNoError(t, ws.WriteMessage(websocket.TextMessage, []byte("3")))
		assertNoError(t, ws.WriteMessage(websocket.TextMessage, []byte(winner)))

		time.Sleep(50 * time.Millisecond)
		assertGameStartedWith(t, game.startedWith, 3)
		assertGameFinishedWith(t, game.finishedWith, winner)

		within(t, 50*time.Millisecond, func() {
			assertWebSocketMsg(t, ws, wantAlert)
		})
	})
}

func assertWebSocketMsg(t testing.TB, ws *websocket.Conn, want string) {
	t.Helper()
	_, msg, err := ws.ReadMessage()
	assertNoError(t, err)
	got := string(msg)
	if got != want {
		t.Errorf("WebSocket message got %q, want %q", got, want)
	}
}

func within(t testing.TB, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{}, 1)

	go func() {
		assert()
		done <- struct{}{}
	}()

	select {
	case <-time.After(d):
		t.Fatal("Timeout")
	case <-done:
	}
}

func mustDialWs(t *testing.T, wsUrl string) *websocket.Conn {
	t.Helper()
	result, _, errDial := websocket.DefaultDialer.Dial(wsUrl, nil)
	assertNoError(t, errDial)
	return result
}

func newGetGameRequest() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/game", nil)
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

func mustMakePlayerServer(t testing.TB, store app.PlayerStore, game app.GameIntf) *app.PlayerServer {
	t.Helper()
	server, err := app.NewPlayerServer(store, game)
	if err != nil {
		t.Fatal("problem creating server", err)
	}
	return server
}

func getLeagueFromResponse(response *httptest.ResponseRecorder) (league app.League, err error) {
	league, err = app.NewLeague(response.Body)
	return
}

func AssertScore(t testing.TB, got, want int) {
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

func assertPlayers(t testing.TB, got app.League, want app.League) {
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
