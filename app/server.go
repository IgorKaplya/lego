package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
	tmpl *template.Template
	game GameIntf
}

func NewPlayerServer(store PlayerStore, game GameIntf) (*PlayerServer, error) {
	result := new(PlayerServer)
	result.store = store
	result.game = game

	tmpl, errParse := template.New("game").Parse(templateGame)
	if errParse != nil {
		return nil, fmt.Errorf("problem parsing game template, %v", errParse)
	}
	result.tmpl = tmpl

	router := http.NewServeMux()
	router.Handle("/players/{name}", http.HandlerFunc(result.playerHandle))
	router.Handle("/league", http.HandlerFunc(result.leagueHandle))
	router.Handle("/game", http.HandlerFunc(result.gameHandle))
	router.Handle("/ws", http.HandlerFunc(result.webSocket))
	result.Handler = router

	return result, nil
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type playerServerWs struct {
	conn *websocket.Conn
}

// Write implements [io.Writer].
func (w *playerServerWs) Write(p []byte) (n int, err error) {
	err = w.conn.WriteMessage(websocket.TextMessage, p)
	n = len(p)
	return
}

func newPlayerServerWs(w http.ResponseWriter, r *http.Request) *playerServerWs {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("problem upgrading connection to web-socket, %v", err)
	}
	return &playerServerWs{conn: conn}
}

func (w *playerServerWs) WaitForMsg() string {
	_, msg, err := w.conn.ReadMessage()
	if err != nil {
		log.Printf("problem reading from websocket, %v", err)
	}
	return string(msg)
}

func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
	ws := newPlayerServerWs(w, r)
	numberOfPlayers, _ := strconv.Atoi(ws.WaitForMsg())
	p.game.Start(numberOfPlayers, ws)

	winner := ws.WaitForMsg()
	p.game.Finish(winner)
}

const templateGame = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Lets play poker</title>
</head>
<body>
<section id="game">
    <div id="game-start">
        <label for="player-count">Number of players</label>
        <input type="number" id="player-count"/>
        <button id="start-game">Start</button>
    </div>

    <div id="declare-winner">
        <label for="winner">Winner</label>
        <input type="text" id="winner"/>
        <button id="winner-button">Declare winner</button>
    </div>

    <div id="blind-value"/>
</section>

<section id="game-end">
    <h1>Another great game of poker everyone!</h1>
    <p><a href="/league">Go check the league table</a></p>
</section>

</body>
<script type="application/javascript">
    const startGame = document.getElementById('game-start')

    const declareWinner = document.getElementById('declare-winner')
    const submitWinnerButton = document.getElementById('winner-button')
    const winnerInput = document.getElementById('winner')

    const blindContainer = document.getElementById('blind-value')

    const gameContainer = document.getElementById('game')
    const gameEndContainer = document.getElementById('game-end')

    declareWinner.hidden = true
    gameEndContainer.hidden = true

    document.getElementById('start-game').addEventListener('click', event => {
        startGame.hidden = true
        declareWinner.hidden = false

        const numberOfPlayers = document.getElementById('player-count').value

        if (window['WebSocket']) {
            const conn = new WebSocket('ws://' + document.location.host + '/ws')

            submitWinnerButton.onclick = event => {
                conn.send(winnerInput.value)
                gameEndContainer.hidden = false
                gameContainer.hidden = true
            }

            conn.onclose = evt => {
                blindContainer.innerText = 'Connection closed'
            }

            conn.onmessage = evt => {
                blindContainer.innerText = evt.data
            }

            conn.onopen = function () {
                conn.send(numberOfPlayers)
            }
        }
    })
</script>
</html>
`

func (p *PlayerServer) gameHandle(w http.ResponseWriter, r *http.Request) {
	p.tmpl.Execute(w, nil)
}

func (p *PlayerServer) playerHandle(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	switch r.Method {
	case http.MethodGet:
		p.processScore(w, name)
	case http.MethodPost:
		p.processWin(w, name)
	}
}

func (p *PlayerServer) leagueHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	league := p.store.GetLeague()
	json.NewEncoder(w).Encode(&league)
}

func (p *PlayerServer) processScore(w http.ResponseWriter, name string) {

	score := p.store.GetPlayerScore(name)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	w.Write([]byte(strconv.Itoa(score)))
}

func (p *PlayerServer) processWin(w http.ResponseWriter, name string) {
	w.WriteHeader(http.StatusAccepted)
	p.store.RecordWin(name)
}

func GetPlayerScore(name string) int {
	if name == "Pepper" {
		return 20
	}

	if name == "Floyd" {
		return 10
	}
	return 0
}
