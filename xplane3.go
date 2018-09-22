package main

import (
	"log"
	"sync"
)

var done chan bool
var wg sync.WaitGroup
var X XPC

func main() {
	var newmsg FlightVal
	newmsg.MsgType = 118
	newmsg.Name = "APHEADING"
	newmsg.Idx = 1
	X.AddMsg(newmsg)
	X.AddMsg(FlightVal{MsgType: 13, Name: "FLAPSSETTING", Idx: 3, Help: "This is the current setting of the flaps handle"})
	X.AddMsg(FlightVal{MsgType: 13, Name: "FLAPSCURRENT", Idx: 4, Help: "This is the current value of the flaps"})
	X.AddMsg(FlightVal{MsgType: 13, Name: "ELEVATORTRIM", Idx: 0, Help: "This is the current setting of the elevator trim"})
	X.AddMsg(FlightVal{MsgType: 13, Name: "AILERONTRIM", Idx: 1, Help: "This is the current setting of the aileron trim"})
	X.AddMsg(FlightVal{MsgType: 13, Name: "RUDDERTRIM", Idx: 2, Help: "This is the current setting of the rudder trim"})

	X.Init("127.0.0.1:49000", ":49003")
	X.Connect()

	go X.Receive()

	server := NewServer("8080")

	log.Println("Starting Server on port %v\n", server.Addr)
	go func() {
		// This starts the HTTPS server using the cert and key locations specified in the
		// config.json file
		//
		log.Println("Starting Server")
		err := server.ListenAndServe()
		if err != nil {
			log.Println("Fatal: Cannot Start Server, exiting ", err)

		}

	}()

	server.WaitShutdown()

	wg.Wait()
	log.Printf("Service Exiting")
}
