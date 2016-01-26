package main

import (
	"fmt"
	"flag"
)

/* standard type for board objects */
type Sprite struct {
	X		int
	Y		int
	Color	rune
}

/* the world in a crystal ball */
var world [][]Sprite


/* Proof of concept program to prove usability of the network design
being intended for Net Rogue */
func main() {
	var gameType, sysType string
	flag.StringVar(&gameType, "game", "checkers", "Choose the game mode [checkers, chess, Go]")
	flag.StringVar(&sysType, "m", "server", "Choose whether to be the server or client [server, client]")
	flag.Parse()
	/* END flag code; begin main method */
	var bheight, bwidth int
	var beginGame func()
	switch gameType {
		case "checkers":
			bheight, bwidth = 9, 9
			beginGame = func() {
				if sysType == "server" {
					m := 'W'
				} else {
					m := 'B'
				}
				beginCheckers(m)
			}
		default:
	}

	// Dynamically allocate the world 2D array
	world = make([][]Sprite, bheight)
	for p, _ := range world {
		world[p] = make([]Sprite, bwidth)
	}

	go beginNetworking()

	go beginGame()

}


