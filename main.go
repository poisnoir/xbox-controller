package main

import (
	"fmt"

	"github.com/holoplot/go-evdev"
)

func main() {
	// Typically /dev/input/eventX. You can find yours using 'evtest'
	// or by listing /dev/input/by-id/
	device, err := evdev.Open("/dev/input/event20")
	if err != nil {
		panic(err)
	}
)

	for {
		ev, err := device.ReadOne()
		if err != nil {
			break
		}
		// Type 1 = Button, Type 3 = Absolute (Sticks/Triggers)
		fmt.Printf("Type: %d, Code: %d, Value: %d\n", ev.Type, ev.Code, ev.Value)
	}
}
