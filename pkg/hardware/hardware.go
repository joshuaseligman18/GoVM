package hardware

import (
	"fmt"

	"github.com/joshuaseligman/GoVM/pkg/util"
)

// Basic hardware struct
type Hardware struct {
	name string // The name of the hardware
	id int // The ID of the hardware
}

// Creates a new hardware
func NewHardware(hwName string, hwId int) *Hardware {
	hw := Hardware { name: hwName, id: hwId }
	return &hw
}

// Logs a message
func (hw *Hardware) Log(msg string) {
	fmt.Printf("[HW: %s %d]: %d - %s\n", hw.name, hw.id, util.GetCurrentTime(), msg)
}

// Gets the name of the hardware
func (hw *Hardware) GetName() string {
	return hw.name
}

// Gets the ID of the hardware
func (hw *Hardware) GetId() int {
	return hw.id
}