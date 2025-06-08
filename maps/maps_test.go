package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("existing key", func(t *testing.T) {
		given := "existing key"
		want := "existing value"
		dic := Dic{given: want}
		got, err := dic.Search(given)
		assertStrings(t, got, want, given)
		assertErrorIs(t, err, nil)
	})

	t.Run("missing key", func(t *testing.T) {
		dic := Dic{}
		given := "missing key"
		got, err := dic.Search(given)
		want := ""
		assertStrings(t, got, want, given)
		assertErrorIs(t, err, ErrKeyNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new key", func(t *testing.T) {
		given := "new key"
		want := "new value"
		dic := Dic{}
		dic.Add(given, want)
		got, err := dic.Search(given)

		assertErrorIs(t, err, nil)
		assertStrings(t, got, want, given)
	})

	t.Run("existing key", func(t *testing.T) {
		given := "xisting key"
		want := "existing value"
		dic := Dic{given: want}
		err := dic.Add(given, want)
		got, _ := dic.Search(given)

		assertErrorIs(t, err, ErrKeyAlreadyExists)
		assertStrings(t, got, want, given)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("new key", func(t *testing.T) {
		given := "new key"
		want := "new value"
		dic := Dic{}
		err := dic.Update(given, want)
		got, _ := dic.Search(given)

		assertErrorIs(t, err, ErrKeyNotFound)
		assertStrings(t, got, "", given)
	})

	t.Run("existing key", func(t *testing.T) {
		given := "existing key"
		want := "updated value"
		dic := Dic{given: "existing value"}
		err := dic.Update(given, want)
		got, _ := dic.Search(given)

		assertErrorIs(t, err, nil)
		assertStrings(t, got, want, given)
	})
}

func TestDelete(t *testing.T) {
	t.Run("missing key", func(t *testing.T) {
		given := "missing key"
		want := ""
		dic := Dic{}
		err := dic.Delete(given)
		got, _ := dic.Search(given)

		assertErrorIs(t, err, ErrKeyNotFound)
		assertStrings(t, got, want, given)
	})

	t.Run("existing key", func(t *testing.T) {
		given := "existing key"
		want := ""
		dic := Dic{given: "existing value"}
		deleteError := dic.Delete(given)
		got, searchError := dic.Search(given)

		assertErrorIs(t, deleteError, nil)
		assertErrorIs(t, searchError, ErrKeyNotFound)
		assertStrings(t, got, want, given)
	})
}

func assertErrorIs(t *testing.T, err error, want error) {
	t.Helper()
	if !errors.Is(err, want) {
		t.Errorf("err %q want %q", err, want)
	}
}

func assertStrings(t *testing.T, got string, want string, given string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, given)
	}
}
