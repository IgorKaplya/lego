package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper time.Duration

func (s DefaultSleeper) Sleep() {
	time.Sleep(time.Duration(s))
}

func Countdown(w io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintf(w, "%d\n", i)
		s.Sleep()
	}
	fmt.Fprint(w, "Go!")
}

func main() {
	Countdown(os.Stdout, DefaultSleeper(3*time.Second))
}
