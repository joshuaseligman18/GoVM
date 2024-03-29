package clock

import (
	"sync"
	"time"

	"github.com/joshuaseligman/GoVM/pkg/hardware"
)

// Struct for tho clock
type Clock struct {
	hw *hardware.Hardware // The hardware struct
	clockListeners []ClockListener // The list of items that listen to the clock
	ticker *time.Ticker // The ticker for each pulse
	running bool // Variable for if the clock is running a program
	wg sync.WaitGroup // Waitgroup to cleanly stop the clock
}

var (
	stopChan chan bool = make(chan bool, 1) // The signal to stop the clock
)

// Creates the clock
func NewClock() *Clock {
	clk := Clock { 
		hw: hardware.NewHardware("CLK", 0),
		clockListeners: make([]ClockListener, 0),
		running: false,
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
	clk.ticker = time.NewTicker(time.Duration(clockTime) * time.Millisecond)
	clk.running = true
	// Run forever
	for {
		select {
		case <-stopChan:
			clk.running = false
			return
		// If the delay is up
		case <-clk.ticker.C:
			clk.wg.Add(1)
			clk.Log("clock pulse initialized")
			// Call the pulse on all the ClockListeners
			for i := 0; i < len(clk.clockListeners); i++ {
				clk.clockListeners[i].Pulse()
			}
			clk.wg.Done()
		}
	}
}

// Starts the clock given the clock time in milliseconds
func (clk *Clock) StartClockAPI(clockTime int, outChan chan []any) {
	// Create the ticker with the given delay
	clk.ticker = time.NewTicker(time.Duration(clockTime) * time.Millisecond)
	clk.running = true
	// Run forever
	for clk.running {
		select {
		case <-stopChan:
			clk.running = false
		// If the delay is up
		case <-clk.ticker.C:
			clk.wg.Add(1)
			clk.Log("clock pulse initialized")
			// Call the pulse on all the ClockListeners
			out := make([]any, len(clk.clockListeners))
			for i := 0; i < len(clk.clockListeners); i++ {
				out[i] = clk.clockListeners[i].Pulse()
			}
			outChan <- out
			clk.wg.Done()
		}
	}
}

// Gets the ticker for the clock
func (clk *Clock) StopClock() {
	// Don't stop until the CPU is ready to stop
	clk.wg.Wait()
	clk.ticker.Stop()
	stopChan <- true
}

// Function to determine if the clock is running
func (clk *Clock) IsStopped() bool {
	return !clk.running
}

// Logs a message
func (clk *Clock) Log(msg string) {
	clk.hw.Log(msg)
}