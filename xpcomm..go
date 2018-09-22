package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"sync"
)

const (
	datagramPrefixLength = 5
	messageLength        = 36
)

type XPC struct {
	RemoteAddress string
	LocalAddress  string
	connection    *net.UDPConn
	FD            FlightDets
}
type FlightVal struct {
	MsgType int
	Name    string
	Help    string
	Idx     int
	Value   float64
	// ADD flags for readable / writable?
}

/*
trim, elev	Elevator trim.
trim, ailrn	Aileron trim.
trim, ruddr	Rudder trim.
flap, handl	•
flap, postn	Flap position.
slat, ratio	Slat position.
sbrak, handl
sbrak, postn	Speedbrake position.
*/
type FlightDets struct {
	Vals    []FlightVal
	fdmutex sync.RWMutex
}

func NewXPC(remoteAddress, localAddress string) XPC {
	return XPC{
		RemoteAddress: remoteAddress,
		LocalAddress:  localAddress,
	}
}

type Command struct {
	Message uint32
	Data    [8]float32
}
type Message interface {
	Type() uint
}

func (x *XPC) Send(command Command) {
	commandData := command.Data

	buf := new(bytes.Buffer)

	buf.Write([]byte{'D', 'A', 'T', 'A', 0})
	buf.Write([]byte{byte(command.Message), 0, 0, 0})

	if err := binary.Write(buf, binary.LittleEndian, &commandData); err != nil {
		fmt.Println(err)
		return
	}

	x.connection.Write(buf.Bytes())
}
func (x *XPC) Init(remoteAddress, localAddress string) {
	x.LocalAddress = localAddress
	x.RemoteAddress = remoteAddress

}
func (x *XPC) AddMsg(newmsg FlightVal) {
	x.FD.fdmutex.Lock()
	defer x.FD.fdmutex.Unlock()
	x.FD.Vals = append(x.FD.Vals, newmsg)
}
func (x *XPC) GetVals() map[string]float64 {
	mp := make(map[string]float64)
	x.FD.fdmutex.RLock()
	defer x.FD.fdmutex.RUnlock()
	for i := 0; i < len(x.FD.Vals); i++ {
		mp[x.FD.Vals[i].Name] = x.FD.Vals[i].Value

	}
	return mp

}

func (x *XPC) GetValue(valname string) (FlightVal, error) {
	x.FD.fdmutex.RLock()
	defer x.FD.fdmutex.RUnlock()
	for i := 0; i < len(x.FD.Vals); i++ {
		if x.FD.Vals[i].Name == valname {
			return x.FD.Vals[i], nil
		}

	}
	var fv FlightVal
	return fv, errors.New("Could not get value")

}
func (x *XPC) Receive() {
	serverAddr, err := net.ResolveUDPAddr("udp", x.LocalAddress)
	serverConn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		panic(err)
	}

	defer serverConn.Close()
	buf := make([]byte, 1024)

	for {
		n, _, _ := serverConn.ReadFromUDP(buf)
		m := (n - datagramPrefixLength) / messageLength

		for i := 0; i < m; i++ {
			sentence := buf[datagramPrefixLength+i*messageLength : datagramPrefixLength+(i+1)*messageLength]
			messageType := sentence[0]
			messageBuffer := bytes.NewBuffer(sentence[4:])

			messageData := make([]float32, 8)
			binary.Read(messageBuffer, binary.LittleEndian, &messageData)
			// Make a copy of the data
			x.FD.fdmutex.RLock()
			workingdata := x.FD.Vals
			x.FD.fdmutex.RUnlock()
			for i := 0; i < len(workingdata); i++ {
				if workingdata[i].MsgType == int(messageType) {
					workingdata[i].Value = float64(messageData[workingdata[i].Idx])
				}

			}
			x.FD.fdmutex.Lock()
			x.FD.Vals = workingdata
			x.FD.fdmutex.Unlock()

		}
	}
}

func (x *XPC) Connect() {
	udpAddr, err := net.ResolveUDPAddr("udp", x.RemoteAddress)

	if err != nil {
		fmt.Println("Wrong address!")
		return
	}

	x.connection, err = net.DialUDP("udp", nil, udpAddr)
}
