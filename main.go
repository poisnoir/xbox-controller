package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/holoplot/go-evdev"
	"github.com/poisnoir/spine-go"
)

func main() {

	namespace := flag.String("namespace", "rime", "spine namespace to join")
	key := flag.String("key", "ppap", "spine namespace key")
	devicePath := flag.String("usb", "/dev/input/event20", "Path to Xbox controller event file")

	flag.Parse()

	device, err := evdev.Open(*devicePath)
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	ns, err := spine.JointNamespace(*namespace, *key, logger)
	if err != nil {
		panic(err)
	}

	pub, err := spine.NewPublisher[XboxController](ns, "xbox-controller")
	if err != nil {
		panic(err)
	}

	var data XboxController

	for {

		ev, err := device.ReadOne()
		if err != nil {
			break
		}
		updateControllerState(&data, ev)

		pub.Publish(data)
	}
}

func updateControllerState(c *XboxController, ev *evdev.InputEvent) {
	switch ev.Type {
	case 1: // EV_KEY (Buttons)
		val := ev.Value == 1 // true if pressed, false if released
		switch ev.Code {
		case 304:
			c.A = val
		case 305:
			c.B = val
		case 307:
			c.X = val
		case 308:
			c.Y = val
		case 310:
			c.LB = val
		case 311:
			c.RB = val
		case 314:
			c.Back = val
		case 315:
			c.Start = val
		case 316:
			c.Guide = val
		case 317:
			c.LSB = val
		case 318:
			c.RSB = val
		}

	case 3: // EV_ABS (Axes, Triggers, D-Pad)
		switch ev.Code {
		// Joysticks (with 10% deadzone approx)
		case 0:
			c.LeftStick.X = applyDeadzone(ev.Value)
		case 1:
			c.LeftStick.Y = applyDeadzone(ev.Value)
		case 3:
			c.RightStick.X = applyDeadzone(ev.Value)
		case 4:
			c.RightStick.Y = applyDeadzone(ev.Value)

		// Triggers
		case 2:
			c.LT = ev.Value
		case 5:
			c.RT = ev.Value

		// D-Pad (Hat Axes)
		case 16: // Hat0X (Left/Right)
			c.Left = ev.Value == -1
			c.Right = ev.Value == 1
		case 17: // Hat0Y (Up/Down)
			c.Up = ev.Value == -1
			c.Down = ev.Value == 1
		}
	}
}

func applyDeadzone(val int32) int32 {
	if val > -2000 && val < 2000 {
		return 0
	}
	return val
}
