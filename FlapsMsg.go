package main

/*
Trim, flaps, slats, and speedbrakes (trim/flap/slat/s-brakes)
-------------------------------------------------------------
0trim, elev	Elevator trim.
1trim, ailrn	Aileron trim.
2trim, ruddr	Rudder trim.
3flap, handl	•
4flap, postn	Flap position.
5slat, ratio	Slat position.
6sbrak, handl
7sbrak, postn	Speedbrake position.
*/
const (
	TrimFlapsSlatsSpeedbrakesMessageType = 13
)

func FlapsCommand(newFlaps float64) Command {
	var cmd Command
	cmd.Message = uint32(TrimFlapsSlatsSpeedbrakesMessageType)
	cmd.Data = [8]float32{-999, -999, -999, float32(newFlaps), -999, -999, -999, -999}
	return cmd
}
