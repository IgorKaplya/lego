package main

import (
	"log"
	"net/http"

	"github.com/IgorKaplya/lego/app"
)

func main() {
	store, errDb, cleanup := app.FileSystemPlayerStoreFromFile("game.db.json")
	if errDb != nil {
		log.Fatalf("problem creating player store, %v", errDb)
	}
	defer cleanup()

	game := app.NewGame(app.BlindAlerterFun(app.Alerter), store)
	server, errServer := app.NewPlayerServer(store, game)
	if errServer != nil {
		log.Fatalf("problem creating server, %v", errServer)
	}

	log.Fatal(http.ListenAndServe(":5000", server))
}
