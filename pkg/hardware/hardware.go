package hardware

type Hardware struct {
	name string
}

func NewHardware(hwName string) *Hardware {
	hw := Hardware { name: hwName }
	return &hw
}

func (hw *Hardware) GetName() string {
	return hw.name
}