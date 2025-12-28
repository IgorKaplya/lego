package app

import "time"

type GameIntf interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

type Game struct {
	alerter BlindAlerter
	store   PlayerStore
}

func NewGame(alerter BlindAlerter, store PlayerStore) *Game {
	return &Game{alerter: alerter, store: store}
}

func (g *Game) Start(numberOfPlayers int) {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Minute
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute
	for _, blindAmount := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blindAmount)
		blindTime = blindTime + blindIncrement
	}
}

func (g *Game) Finish(winner string) {
	g.store.RecordWin(winner)
}
