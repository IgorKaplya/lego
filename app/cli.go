package app

import (
	"bufio"
	"io"
	"strings"
)

type Cli struct {
	store   PlayerStore
	scanner *bufio.Scanner
}

func NewCli(store PlayerStore, in io.Reader) *Cli {
	return &Cli{
		store:   store,
		scanner: bufio.NewScanner(in),
	}
}

func (c *Cli) PlayPoker() {
	c.scanner.Scan()
	player := strings.Replace(c.scanner.Text(), " wins", "", 1)
	c.store.RecordWin(player)
}
