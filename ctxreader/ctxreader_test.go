package ctxreader_test

import (
	"context"
	"strings"
	"testing"

	"github.com/IgorKaplya/lego/ctxreader"
)

func TestContextAwareReader(t *testing.T) {
	t.Run("Normal reader", func(t *testing.T) {
		reader := strings.NewReader("123456")
		got := make([]byte, 3)
		if _, err := reader.Read(got); err != nil {
			t.Fatalf("unexpected error 1, %s", err)
		}
		assertBufferHas(t, got, "123")

		if _, err := reader.Read(got); err != nil {
			t.Fatalf("unexpected error 2, %s", err)
		}
		assertBufferHas(t, got, "456")

	})
	t.Run("Cancellable reader", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		reader := ctxreader.NewCancellableReader(strings.NewReader("123456"), ctx)

		got := make([]byte, 3)
		if _, err := reader.Read(got); err != nil {
			t.Fatalf("unexpected error 1, %s", err)
		}
		assertBufferHas(t, got, "123")

		cancel()
		n, err := reader.Read(got)
		if err == nil {
			t.Fatal("expected cancel error")
		}
		assertRead(t, n, 0)
	})
}

func assertRead(t *testing.T, n, want int) {
	if n != want {
		t.Errorf("read %d, want %d", n, want)
	}
}

func assertBufferHas(t testing.TB, buf []byte, want string) {
	t.Helper()
	got := string(buf)
	if want != got {
		t.Errorf("Buffer got %q, want %q", got, want)
	}
}
