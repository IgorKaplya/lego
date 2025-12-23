package main

import (
	"fmt"
	"log"
	"os"

	"github.com/IgorKaplya/lego/app"
)

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	store, err, cleanup := app.FileSystemPlayerStoreFromFile("game.db.json")
	if err != nil {
		log.Fatalf("problem creating player store, %v", err)
	}
	defer cleanup()

	cli := app.NewCli(store, os.Stdin)
	cli.PlayPoker()
}
