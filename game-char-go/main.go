package main

import (
	"github.com/achmadsy/game-char-go/db"
	"github.com/achmadsy/game-char-go/routes"
)

func main() {
	db.InitDB()
	routes.Run()
}
