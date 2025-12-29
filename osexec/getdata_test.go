package osexec_test

import (
	"strings"
	"testing"

	"github.com/IgorKaplya/lego/osexec"
)

func TestGetData(t *testing.T) {
	reader := strings.NewReader(`
<payload>
    <message>aRrGh</message>
</payload>`)

	got := osexec.GetData(reader)

	assertGetData(t, got, "ARRGH")
}

func assertGetData(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("GetData got %q, want %q", got, want)
	}
}
