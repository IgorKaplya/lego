package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("single thread", func(t *testing.T) {
		want := 3
		counter := NewCounter()

		for range want {
			counter.Inc()
		}

		assertCount(t, counter, want)
	})
	t.Run("multi thread", func(t *testing.T) {
		want := 1000
		counter := NewCounter()

		var waitGroup sync.WaitGroup
		waitGroup.Add(1000)
		for range want {
			go func() {
				counter.Inc()
				waitGroup.Done()
			}()
		}
		waitGroup.Wait()

		assertCount(t, counter, want)
	})
}

func assertCount(t testing.TB, counter *Counter, want int) {
	got := counter.Value()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
