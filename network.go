package main

import (
	"fmt"
	"net"
	"time"
	"encoding/json"
)

/***
Network communications format:
	client -> server: "QUERY"
	server -> client:
		"UPDATE" - there is a world update incoming
		"WAIT"   - wait for an update
	client -> server: 
		"OK" - acknowledge preparedness for an update
		"WAITING" - more or less ignore this, placeholder, waiting ensues
	if we have an update...
		server -> client: [JSON-encoded world]
	else...
		~nothing happens~; aka go on to next cycle
***/


/* sends a stream of json over the network for transferring the world across the net */
func sendWorld(c net.Conn) {
	enc := json.NewEncoder(c)
	err := enc.Encode(world)
	check(err)
}

/* initializes the client connection */
func beginClient(srvAddr string) {
	conn, err := net.Dial("tcp", srvAddr)
	checkHard(err)

	//communication goes here 
}

/* initializes the primary network framework */
func beginServer() {
	ln, err := net.Listen("tcp", ":5573")
	checkHard(err)

	go acceptor(ln)
}

/* accepts and passes on connections */
func acceptor(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		check(err)
		go handler(conn)
		time.Sleep(10 * time.Millisecond)
	}
}

/* handles incoming connections and keeps them in sync */
func handler(conn net.Conn) {
	for {
		//manage the player number, game, states, etc for the connection

		time.Sleep(10 * time.Millisecond)
	}
}



