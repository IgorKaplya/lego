package app

import (
	"bytes"
	"testing"
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

func newDatabase(data string) FileDatabase {
	return &TestDatabase{data: data}
}

func TestFileSystemStore(t *testing.T) {
	store, err := NewFileSystemPlayerStore(newDatabase(`[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}
	]`))
	AssertNoError(t, err)

	t.Run("league from reader", func(t *testing.T) {

		got := store.GetLeague()
		got = store.GetLeague()

		want := League{
			{Name: "Chris", Wins: 33},
			{Name: "Cleo", Wins: 10},
		}
		AssertPlayers(t, got, want)
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
