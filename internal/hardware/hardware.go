package hardware

// Basic hardware struct
type Hardware struct {
	name string
	id int
}

// Creates a new hardware
func NewHardware(hwName string, hwId int) *Hardware {
	hw := Hardware { name: hwName, id: hwId }
	return &hw
}

// Gets the name of the hardware
func (hw *Hardware) GetName() string {
	return hw.name
}

// Gets the ID of the hardware
func (hw *Hardware) GetId() int {
	return hw.id
}