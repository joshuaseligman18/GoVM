package clock

// Interface that listen to the clock
type ClockListener interface {
	Pulse() any // Function that gets called every clock cycle
}