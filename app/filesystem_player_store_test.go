package app_test

import (
	"bytes"
	"testing"

	"github.com/IgorKaplya/lego/app"
)

type TestDatabase struct {
	data string
}

// Truncate implements [FileDatabase].
func (t *TestDatabase) Truncate(size int64) error {
	t.data = ""
	return nil
}

// Read implements [FileDatabase].
func (t *TestDatabase) Read(p []byte) (n int, err error) {
	buffer := bytes.NewBufferString(t.data)
	return buffer.Read(p)
}

// Seek implements [FileDatabase].
func (t *TestDatabase) Seek(offset int64, whence int) (n int64, err error) {
	return
}

// Write implements [FileDatabase].
func (t *TestDatabase) Write(p []byte) (n int, err error) {
	buffer := bytes.Buffer{}
	n, err = buffer.Write(p)
	t.data = buffer.String()
	return
}

func newDatabase(data string) app.FileDatabase {
	return &TestDatabase{data: data}
}

func TestFileSystemStore(t *testing.T) {
	store, err := app.NewFileSystemPlayerStore(newDatabase(`[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}
	]`))
	assertNoError(t, err)

	t.Run("league from reader", func(t *testing.T) {

		got := store.GetLeague()
		got = store.GetLeague()

		want := app.League{
			{Name: "Chris", Wins: 33},
			{Name: "Cleo", Wins: 10},
		}
		assertPlayers(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		got := store.GetPlayerScore("Chris")
		want := 33
		AssertScore(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		want := 34
		AssertScore(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		store.RecordWin("Pepper")
		got := store.GetPlayerScore("Pepper")
		want := 1
		AssertScore(t, got, want)
	})
}
