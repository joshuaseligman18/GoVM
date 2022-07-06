package clock

import (
	"time"
	"github.com/joshuaseligman/GoVM/pkg/hardware"
)

// Struct for tho clock
type Clock struct {
	hw *hardware.Hardware
	clockListeners []ClockListener
}

// Creates the clock
func NewClock() *Clock {
	clk := Clock { 
		hw: hardware.NewHardware("CLK", 0),
		clockListeners: make([]ClockListener, 0),
	}
	return &clk
}

// Adds a ClockListener to be called every clock cycle
func (clk *Clock) AddClockListener(clockListener ClockListener) {
	clk.clockListeners = append(clk.clockListeners, clockListener)
}

// Starts the clock given the clock time in milliseconds
func (clk *Clock) StartClock(clockTime int) {
	// Create the ticker with the given delay
	ticker := time.NewTicker(time.Duration(clockTime) * time.Millisecond)
	// Run forever
	for {
		select {
		// If the delay is up
		case <-ticker.C:
			clk.Log("clock pulse initialized")
			// Call the pulse on all the ClockListeners
			for i := 0; i < len(clk.clockListeners); i++ {
				clk.clockListeners[i].Pulse()
			}
		}
	}
}

// Logs a message
func (clk *Clock) Log(msg string) {
	clk.hw.Log(msg)
}