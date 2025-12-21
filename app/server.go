package main

import (
	"net/http"
	"strconv"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/players/")
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
