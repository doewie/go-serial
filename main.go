package main

import (
	"fmt"
	"os"
)

var (
	serialRecv = make(chan []byte)
	serialSend = make(chan []byte)
	done       = make(chan error)
)

func main() {
	// printPorts()
	// printPortDetails()
	openPortCom5()

	go SerialDispatch()

	if err, ok := <-done; ok {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}

func check(e error) {
	if e != nil {
		done <- e
	}
}
