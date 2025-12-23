package app_test

import (
	"strings"
	"testing"

	"github.com/IgorKaplya/lego/app"
)

func TestCli(t *testing.T) {
	t.Run("chris wins", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := &app.StubPlayerStore{}
		cli := app.NewCli(store, in)

		cli.PlayPoker()

		app.AssertWinCalls(t, store.WinCalls, []string{"Chris"})
	})

	t.Run("cleo wins", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		store := &app.StubPlayerStore{}
		cli := app.NewCli(store, in)

		cli.PlayPoker()

		app.AssertWinCalls(t, store.WinCalls, []string{"Cleo"})
	})
}
