package main

import (
	"log"
	"net/http"
	"os"

	"github.com/IgorKaplya/lego/app"
)

const dbFileName = "game.db.json"

func main() {
	file, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening file %q %v", dbFileName, err)
	}

	store, err := app.NewFileSystemPlayerStore(file)
	if err != nil {
		log.Fatalf("problem creating store, %v", err)
	}

	server := app.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))
}
