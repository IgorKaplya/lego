package main

import (
	"bytes"
	"reflect"
	"testing"
)

type SleeperStub struct {
}

func (s *SleeperStub) Sleep() {
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")
	return 0, nil
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

func TestCountdown(t *testing.T) {
	t.Run("output is correct", func(t *testing.T) {
		wantPrinted := "3\n2\n1\nGo!"
		var writer = new(bytes.Buffer)
		var sleeper = new(SleeperStub)

		Countdown(writer, sleeper)

		if wantPrinted != writer.String() {
			t.Errorf("got %q want %q", writer, wantPrinted)
		}
	})
	t.Run("write and sleep sequence is correct", func(t *testing.T) {
		want := []string{"write", "sleep", "write", "sleep", "write", "sleep", "write"}
		spy := new(CountdownOperationsSpy)

		Countdown(spy, spy)

		if !reflect.DeepEqual(want, spy.Calls) {
			t.Errorf("got %v want %v", spy.Calls, want)
		}
	})
}
