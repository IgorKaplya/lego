package main

import (
	"net/http"
	"strconv"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	result := &PlayerServer{}
	result.store = store
	router := http.NewServeMux()
	router.Handle("/players/{name}", http.HandlerFunc(result.playerHandle))
	result.Handler = router

	return result
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
