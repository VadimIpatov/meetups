package main

import (
	"machine"

	"tinygo.org/x/drivers/buzzer"
)

type note struct {
	tone     float64
	duration float64
}

func main() {
	bzrPin := machine.D11
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	bzr := buzzer.New(bzrPin)

	song := []note{
		{buzzer.D4, buzzer.Quarter},
		{buzzer.G4, buzzer.Eighth / 2},
		{buzzer.C5, buzzer.Quarter + buzzer.Eighth},
		{buzzer.B4, buzzer.Eighth},
		{buzzer.G4, buzzer.Eighth/2 + buzzer.Eighth},
		{buzzer.E4, buzzer.Eighth/2 + buzzer.Eighth},
		{buzzer.A4, buzzer.Eighth/2 + buzzer.Eighth},
		{buzzer.D5, buzzer.Half},
	}

	for _, val := range song {
		bzr.Tone(val.tone, val.duration)
	}
}
