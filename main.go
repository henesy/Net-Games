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

/* max size of communications over the network */
var maxCommLen int


/* Proof of concept program to prove usability of the network design
being intended for Net Rogue */
func main() {
	var gameType, sysType, srvAddr string
	flag.StringVar(&gameType, "game", "checkers", "Choose the game mode [checkers, chess, Go]")
	flag.StringVar(&sysType, "m", "server", "Choose whether to be the server or client [server, client]")
	flag.StringVar(&srvAddr, "s", "127.0.0.1:5573", "Used in conjunction with `-m client`: defines server IP:Port address [127.0.0.1:5573]")
	flag.Parse()
	/* END flag code; begin main method */
	var bheight, bwidth int
	var beginGame, beginNetworking func()

	//customize game type and networking
	switch gameType {
		case "checkers":
			plyr := '哑'
			bheight, bwidth = 9, 9
			beginGame = func() {
				if sysType == "server" {
					plyr = 'W'
				} else {
					plyr = 'B'
				}
				beginCheckers(plyr)
			}
		default:
	}
	switch sysType {
		case "server":
			beginNetworking = func() {
				beginServer()
			}
		case "client":
			beginNetworking = func() {
				beginClient(srvAddr)
			}
	}

	//dynamically allocate the world 2D array
	world = make([][]Sprite, bheight)
	for p, _ := range world {
		world[p] = make([]Sprite, bwidth)
	}
	//dynamically set max communications size
	eW := encodeWorld()
	maxCommLen = len(eW)

	go beginNetworking()

	go beginGame()

}


