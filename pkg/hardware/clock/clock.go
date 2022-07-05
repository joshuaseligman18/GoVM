package clock

import "github.com/joshuaseligman/GoVM/pkg/hardware"

// Struct for tho clock
type Clock struct {
	hw *hardware.Hardware
}

// Creates the clock
func NewClock() *Clock {
	clk := Clock { hw: hardware.NewHardware("CLK", 0) }
	return &clk
}

// Logs a message
func (clk *Clock) Log(msg string) {
	clk.hw.Log(msg)
}