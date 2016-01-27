package main

import (
	"fmt"
	"os"
	"encoding/json"
)


/* basic error checker, prints to standard output */
func check(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

/* The less than gentle manner of handling critical (system) errors */
func checkHard(err error) {
	check(err)
	os.Exit(1)
}

/* encodes the world in JSON and returns said world */
func encodeWorld() []byte {
	encW, err := json.Marshal(world)
	check(err)

	return encW
}


