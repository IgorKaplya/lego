package app_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/IgorKaplya/lego/app"
)

type GameSpy struct {
	started      bool
	startedWith  int
	finishedWith string
	blindAlert   []byte
}

// Finish implements [app.GameIntf].
func (g *GameSpy) Finish(winner string) {
	g.finishedWith = winner
}

// Start implements [app.GameIntf].
func (g *GameSpy) Start(numberOfPlayers int, to io.Writer) {
	g.startedWith = numberOfPlayers
	g.started = true
	to.Write(g.blindAlert)
}

func TestPlayPoker(t *testing.T) {

	t.Run("plays", func(t *testing.T) {
		in := strings.NewReader("5\nbobo wins")
		out := new(bytes.Buffer)
		game := &GameSpy{}
		cli := app.NewCli(in, out, game)

		cli.PlayPoker()

		assertPromptText(t, out.String(), app.PlayerPrompt)
		assertGameStartedWith(t, game.startedWith, 5)
		assertGameFinishedWith(t, game.finishedWith, "bobo")
	})

	t.Run("errors on NaN for user num", func(t *testing.T) {
		in := strings.NewReader("Woop\n")
		out := new(bytes.Buffer)
		game := &GameSpy{}
		cli := app.NewCli(in, out, game)

		cli.PlayPoker()

		assertGameStarted(t, game.started, false)
		assertPromptText(t, out.String(), app.PlayerPrompt, app.NaNErrorMessage)
	})

	t.Run("errors on wrong win pattern", func(t *testing.T) {
		in := strings.NewReader("3\nlabadabadab dab")
		out := new(bytes.Buffer)
		game := &GameSpy{}
		cli := app.NewCli(in, out, game)

		cli.PlayPoker()

		assertPromptText(t, out.String(), app.PlayerPrompt, app.WrongWinPatternMessage)
	})
}

func assertGameFinishedWith(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Winner got %q, want %q", got, want)
	}
}

func assertGameStarted(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Fatalf("game started got %v, want %v", got, want)
	}
}

func assertGameStartedWith(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Playert number got %d, want %d", got, want)
	}
}

func assertPromptText(t testing.TB, got string, wantMessages ...string) {
	t.Helper()
	want := strings.Join(wantMessages, "")

	if got != want {
		t.Errorf("Prompt got %q, want %q", got, want)
	}
}
