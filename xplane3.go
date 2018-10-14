package main

import (
	"log"
	"sync"
)

var done chan bool
var wg sync.WaitGroup
var X XPC

func main() {
	done = make(chan bool)

	/* 17
	   Pitch, roll, headings
	   pitch, deg	The aircraft’s pitch, measured in body-axis Euler angles.
	   roll, deg	The aircraft’s roll, measured in body-axis Euler angles.
	   hding, true	The aircraft’s true heading, measured in body-axis Euler angles.
	   hding, mag	The aircraft’s magneti
	*/
	X.AddMsg(FlightVal{MsgType: 17, Name: "PITCH", Idx: 0, Help: "Pitch"})
	X.AddMsg(FlightVal{MsgType: 17, Name: "ROLL", Idx: 1, Help: "Roll"})
	X.AddMsg(FlightVal{MsgType: 17, Name: "HdgTrue", Idx: 2, Help: "Heading - true"})
	X.AddMsg(FlightVal{MsgType: 17, Name: "HdgMag", Idx: 3, Help: "Heading - magnetic"})

	X.AddMsg(FlightVal{MsgType: 118, Name: "APVERTICALVELOCITY", Idx: 2, Help: "This is the Vertcal Vel for the AP"})
	/* 108
	   Autopilot, flight director, and head-up display (switches 3: AP/f-dir/HUD)
	   ap, src	•
	   fdir, mode	•
	   fdir, ptch	•
	   fdir, roll	•
	   HUD, power	•
	   HUD, brite	•
	*/
	X.AddMsg(FlightVal{MsgType: 108, Name: "APSRC", Idx: 0, Help: "AP SRC"})

	/* 116
	   Armed autopilot functions (autopilot arms)
	   nav, arm	•
	   alt, arm	•
	   app, arm	•
	   vnav, enab	•
	   vnav, arm	•
	   vnav, time	•
	   gp, enabl	•
	   app, typ	•
	*/
	X.AddMsg(FlightVal{MsgType: 116, Name: "APNAVARM", Idx: 0, Help: "AP SRC"})
	X.AddMsg(FlightVal{MsgType: 116, Name: "APALTARM", Idx: 1, Help: "AP SRC"})
	X.AddMsg(FlightVal{MsgType: 116, Name: "APAPPARM", Idx: 2, Help: "AP SRC"})
	X.AddMsg(FlightVal{MsgType: 116, Name: "APVNAVENAB", Idx: 3, Help: "AP SRC"})
	X.AddMsg(FlightVal{MsgType: 116, Name: "APVNAVARM", Idx: 4, Help: "AP SRC"})
	X.AddMsg(FlightVal{MsgType: 116, Name: "APVNAVTIME", Idx: 5, Help: "AP SRC"})
	X.AddMsg(FlightVal{MsgType: 116, Name: "APNGPENAB", Idx: 6, Help: "AP SRC"})
	X.AddMsg(FlightVal{MsgType: 116, Name: "APAPPTYP", Idx: 7, Help: "AP SRC"})

	/*
	   98

	   NAV 1/2 OBS
	   NAV n, OBS	•
	   NAV n, s-crs	•
	   NAV n, flag	•
	*/
	X.AddMsg(FlightVal{MsgType: 98, Name: "NAV1OBS", Idx: 0, Help: "Nav1 Obs"})
	X.AddMsg(FlightVal{MsgType: 98, Name: "NAV1SCRS", Idx: 1, Help: "Nav1 s-crs"})
	X.AddMsg(FlightVal{MsgType: 98, Name: "NAV1FLAG", Idx: 2, Help: "Nav1 FLAG"})
	X.AddMsg(FlightVal{MsgType: 98, Name: "NAV2OBS", Idx: 3, Help: "Nav2 Obs"})
	X.AddMsg(FlightVal{MsgType: 98, Name: "NAV2SCRS", Idx: 4, Help: "Nav2 s-crs"})
	X.AddMsg(FlightVal{MsgType: 98, Name: "NAV2FLAG", Idx: 5, Help: "Nav2 FLAG"})

	/*
	   Nav 1 - 99
	   Nav 2 - 100

	   NAV n deflections
	   NAV n, n-typ	•
	   NAV n, to-fr	•
	   NAV n, m-crs	•
	   NAV n, r-brg	•
	   NAV n, dme-d	•
	   NAV n, h-def	•
	   NAV n, v-def	•

	*/
	X.AddMsg(FlightVal{MsgType: 99, Name: "NAV1NTYP", Idx: 0, Help: "Nav1 n-typ"})
	X.AddMsg(FlightVal{MsgType: 99, Name: "NAV1TOFR", Idx: 1, Help: "Nav1 to-fr"})
	X.AddMsg(FlightVal{MsgType: 99, Name: "NAV1MCRS", Idx: 2, Help: "Nav1 m-crs"})
	X.AddMsg(FlightVal{MsgType: 99, Name: "NAV1RBRG", Idx: 3, Help: "Nav1 RBRG"})
	X.AddMsg(FlightVal{MsgType: 99, Name: "NAV1DMED", Idx: 4, Help: "Nav1 dme-d", Logtoconsole: true})
	X.AddMsg(FlightVal{MsgType: 99, Name: "NAV1HDEF", Idx: 5, Help: "Nav1 h-def"})
	X.AddMsg(FlightVal{MsgType: 99, Name: "NAV1VDEF", Idx: 6, Help: "Nav1 v-def"})

	X.AddMsg(FlightVal{MsgType: 100, Name: "NAV2NTYP", Idx: 0, Help: "Nav2 n-typ"})
	X.AddMsg(FlightVal{MsgType: 100, Name: "NAV2TOFR", Idx: 1, Help: "Nav2 to-fr"})
	X.AddMsg(FlightVal{MsgType: 100, Name: "NAV2MCRS", Idx: 2, Help: "Nav2 m-crs"})
	X.AddMsg(FlightVal{MsgType: 100, Name: "NAV2RBRG", Idx: 3, Help: "Nav2 RBRG"})
	X.AddMsg(FlightVal{MsgType: 100, Name: "NAV2DMED", Idx: 4, Help: "Nav2 dme-d"})
	X.AddMsg(FlightVal{MsgType: 100, Name: "NAV2HDEF", Idx: 5, Help: "Nav2 h-def"})
	X.AddMsg(FlightVal{MsgType: 100, Name: "NAV2VDEF", Idx: 6, Help: "Nav2 v-def"})

	X.AddMsg(FlightVal{MsgType: 118, Name: "APHEADING", Idx: 1, Help: "This is the heading for the AP"})
	X.AddMsg(FlightVal{MsgType: 118, Name: "APVERTICALVELOCITY", Idx: 2, Help: "This is the Vertcal Vel for the AP"})

	X.AddMsg(FlightVal{MsgType: 118, Name: "APALT", Idx: 3, Help: "This is the alt for the AP"})
	X.AddMsg(FlightVal{MsgType: 118, Name: "APSPEED", Idx: 0, Help: "This is the alt for the AP"})
	// Autopilot Modes - 117
	X.AddMsg(FlightVal{MsgType: 117, Name: "APAUTOTHROTTLEMODE", Idx: 0, Help: "Autothrottle mode"})
	X.AddMsg(FlightVal{MsgType: 117, Name: "APHEADINGMODE", Idx: 1, Help: "Heading mode"})
	X.AddMsg(FlightVal{MsgType: 117, Name: "APALTMODE", Idx: 2, Help: "Altitude mode"})
	X.AddMsg(FlightVal{MsgType: 117, Name: "APBACMODE", Idx: 3, Help: "BAC mode"})
	X.AddMsg(FlightVal{MsgType: 117, Name: "APAPPMODE", Idx: 4, Help: "APP mode"})
	X.AddMsg(FlightVal{MsgType: 117, Name: "APSYNC", Idx: 5, Help: "APSYNC mode"})

	X.AddMsg(FlightVal{MsgType: 13, Name: "FLAPSSETTING", Idx: 3, Help: "This is the current setting of the flaps handle"})
	X.AddMsg(FlightVal{MsgType: 13, Name: "FLAPSCURRENT", Idx: 4, Help: "This is the current value of the flaps"})
	X.AddMsg(FlightVal{MsgType: 13, Name: "ELEVATORTRIM", Idx: 0, Help: "This is the current setting of the elevator trim"})
	X.AddMsg(FlightVal{MsgType: 13, Name: "AILERONTRIM", Idx: 1, Help: "This is the current setting of the aileron trim"})
	X.AddMsg(FlightVal{MsgType: 13, Name: "RUDDERTRIM", Idx: 2, Help: "This is the current setting of the rudder trim"})
	X.AddMsg(FlightVal{MsgType: 14, Name: "GEAR", Idx: 0, Help: "This is the current setting of the landing gear"})
	X.AddMsg(FlightVal{MsgType: 14, Name: "BRAKES", Idx: 1, Help: "This is the current setting of the Brakes"})
	X.AddMsg(FlightVal{MsgType: 67, Name: "GEARDEPLOYMENT", Idx: 1, Help: "This is the current state of the landing gear"})
	X.AddMsg(FlightVal{MsgType: 67, Name: "GEARFORCE", Idx: 0, Help: "This is the current value for the force on the gear"})
	X.AddMsg(FlightVal{MsgType: 127, Name: "GEARWARNING", Idx: 5, Help: "This is gear warning"})
	X.AddMsg(FlightVal{MsgType: 127, Name: "GEARWORK", Idx: 4, Help: "This is the gear working warning"})
	X.AddMsg(FlightVal{MsgType: 97, Name: "NAV1CurrentFrequency", Idx: 0, Help: "Current Frequency for Nav1"})
	X.AddMsg(FlightVal{MsgType: 97, Name: "NAV1StandbyFrequency", Idx: 1, Help: "Standby Frequency for Nav1"})
	X.AddMsg(FlightVal{MsgType: 97, Name: "NAV2CurrentFrequency", Idx: 4, Help: "Current Frequency for Nav2"})
	X.AddMsg(FlightVal{MsgType: 97, Name: "NAV2StandbyFrequency", Idx: 5, Help: "Standby Frequency for Nav2"})

	X.Init("127.0.0.1:49000", ":49003")
	X.Connect()

	go X.Receive()

	server := NewServer("8080")

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
