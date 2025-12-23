package app

import (
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   League
}

func (s *StubPlayerStore) GetLeague() (result League) {
	return s.League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func GetLeagueFromResponse(response *httptest.ResponseRecorder) (league League, err error) {
	league, err = NewLeague(response.Body)
	return
}

func AssertScore(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("score got %d, want %d", got, want)
	}
}

func AssertContentIsJson(t testing.TB, response *httptest.ResponseRecorder) {
	got := response.Header().Get("content-type")
	want := "application/json"

	if got != want {
		t.Fatalf("content type got %q, want %q", got, want)
	}
}

func AssertPlayers(t testing.TB, got League, want League) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("league got %v want %v", got, want)
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("did not expect an error: %#v", err)
	}
}

func AssertWinCalls(t testing.TB, got []string, want []string) {
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

func AssertStatusCode(t *testing.T, response *httptest.ResponseRecorder, want int) {
	t.Helper()

	got := response.Code

	if got != want {
		t.Errorf("got status %d want %d", got, want)
	}
}

func AssertResponseBody(t *testing.T, response *httptest.ResponseRecorder, want string) {
	got := response.Body.String()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
