package app_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/IgorKaplya/lego/app"
)

type spyAlert struct {
	duration time.Duration
	amount   int
}

func (a spyAlert) String() string {
	return fmt.Sprintf("%d chips at %v", a.amount, a.duration)
}

type SpyBlindAlerter struct {
	alerts []spyAlert
}

// ScheduleAlertAt implements [app.BlindAlerter].
func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, struct {
		duration time.Duration
		amount   int
	}{
		duration: duration,
		amount:   amount,
	})
}

func TestStart(t *testing.T) {
	dummyStore := new(app.StubPlayerStore)
	alerter := new(SpyBlindAlerter)

	game := app.NewGame(alerter, dummyStore)
	game.Start(7)

	wantAlerts := []spyAlert{
		{duration: 0 * time.Minute, amount: 100},
		{duration: 12 * time.Minute, amount: 200},
		{duration: 24 * time.Minute, amount: 300},
		{duration: 36 * time.Minute, amount: 400},
		{duration: 48 * time.Minute, amount: 500},
		{duration: 60 * time.Minute, amount: 600},
		{duration: 72 * time.Minute, amount: 800},
		{duration: 84 * time.Minute, amount: 1000},
		{duration: 96 * time.Minute, amount: 2000},
		{duration: 108 * time.Minute, amount: 4000},
		{duration: 120 * time.Minute, amount: 8000},
	}

	assertAlerts(t, alerter.alerts, wantAlerts)
}

func TestFinisg(t *testing.T) {
	store := new(app.StubPlayerStore)
	dummyAlerter := new(SpyBlindAlerter)
	game := app.NewGame(dummyAlerter, store)

	game.Finish("Maik")

	app.AssertWinCalls(t, store.WinCalls, []string{"Maik"})
}

func assertAlerts(t *testing.T, gotAlerts []spyAlert, wantAlerts []spyAlert) {
	t.Helper()
	gotLen := len(gotAlerts)
	wantLen := len(wantAlerts)
	if gotLen != wantLen {
		t.Fatalf("alerts len got %d want %d", gotLen, wantLen)
	}

	for i, want := range wantAlerts {
		t.Run(want.String(), func(t *testing.T) {
			t.Helper()
			got := gotAlerts[i]
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
