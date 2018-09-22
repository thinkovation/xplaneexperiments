package main

// [0 10 0 160 0 33.878597 0.059031606 -0.012948804]
/*
set, speed	The desired speed setting for the autopilot.
set, hding	The desired heading setting for the autopilot.
set, vvi	The desired vertical velocity setting for the autopilot.
dial, alt	•
vnav, alt	•
use, alt	•
sync, roll	•
sync, pitch	•
*/

const (
	APValuesMessageType = 118
)

func HDGCommand(newheading float64) Command {
	var cmd Command
	cmd.Message = uint32(APValuesMessageType)
	cmd.Data = [8]float32{-999, float32(newheading), -999, -999, -999, -999, -999, -999}
	return cmd
}
