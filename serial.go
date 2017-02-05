package main

import (
	"fmt"
	"log"
	//	"time"

	"go.bug.st/serial.v1"
	"go.bug.st/serial.v1/enumerator"
)

// print serial ports

var activePort serial.Port

func printPorts() {

	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
}

func printPortDetails() {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		fmt.Println("No serial ports found!")
		return
	}
	for _, port := range ports {
		fmt.Printf("Found port: %s\n", port.Name)
		if port.IsUSB {
			fmt.Printf("   USB ID     %s:%s\n", port.VID, port.PID)
			fmt.Printf("   USB serial %s\n", port.SerialNumber)
		}
	}
}

// open a fixed port
// here open COM5
func openPortCom5() {
	// ports, _ := serial.GetPortsList()
	var err error
	activePort, err = serial.Open(
		"COM5",
		&serial.Mode{
			BaudRate: 57600,
			Parity:   serial.NoParity,
			DataBits: 8,
			StopBits: serial.OneStopBit,
		})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to COM5\n")

}

// SerialDispatch handles all incoming and outgoing serial data.
func SerialDispatch() {
	//	go func() {
	//		for data := range serialSend {
	//			activePort.Write(data)
	//		}
	//	}()

	go func() {
		d := make([]byte, 100)
		for {
			n, err := activePort.Read(d)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				break
			}
			if n > 0 {
				fmt.Printf("Received from serial %d characters: *%v* = *%s*\n", n, d[:n], d[:n])
				serialRecv <- d[:n]
			}
			// time.Sleep(1000 * time.Millisecond)
		}
	}()

	for {
		select {

		case data, ok := <-serialRecv:
			if !ok {
				return
			}
			fmt.Printf("received: %s\n", data)
		}
	}
}
