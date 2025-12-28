package app

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Cli struct {
	scanner *bufio.Scanner
	out     io.Writer
	game    GameIntf
}

const PlayerPrompt = "Please enter the number of players:"
const NaNErrorMessage = "Not a number"
const WrongWinPatternMessage = "Wrong pattern, should be: {Name} wins"

func NewCli(in io.Reader, out io.Writer, game GameIntf) *Cli {
	return &Cli{
		scanner: bufio.NewScanner(in),
		out:     out,
		game:    game,
	}
}

func (c *Cli) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)
	c.scanner.Scan()
	numberOfPlayers, err := strconv.Atoi(c.scanner.Text())
	if err != nil {
		fmt.Fprint(c.out, NaNErrorMessage)
		return
	}

	c.game.Start(numberOfPlayers)

	c.scanner.Scan()
	input := c.scanner.Text()
	player := strings.Replace(input, " wins", "", 1)
	if player == input {
		fmt.Fprint(c.out, WrongWinPatternMessage)
	}

	c.game.Finish(player)
}
