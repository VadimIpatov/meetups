package main

import (
	"machine"
	"time"
)

func main() {
	println("Hello 1NCE team!")

	led := machine.D53
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for i := 0; i < 10; i++ {
		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}
