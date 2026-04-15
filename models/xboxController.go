package models

type XboxController struct {
	// Face Buttons
	A, B, X, Y bool

	// Bumpers
	LB, RB bool

	LT, RT int32

	// Joysticks
	LeftStick  Joystick
	RightStick Joystick

	// Thumbstick Clicks
	LSB, RSB bool

	// D-Pad
	Up, Down, Left, Right bool

	// Menu Buttons
	Back, Start, Guide bool
}

type Joystick struct {
	X int32 // Left is negative, Right is positive
	Y int32 // Up is negative, Down is positive
}
