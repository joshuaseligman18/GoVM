package gui

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/joshuaseligman/GoVM/pkg/hardware/cpu"
	"github.com/joshuaseligman/GoVM/pkg/util"
)

// Struct for the data represented within the GUI
type GuiData struct {
	cpu *cpu.Cpu // The CPU for making updates
	timeLabel *widget.Label // The label for the time
	curTime *widget.Label // The current time
	pcLabel *widget.Label // The label for the program counter
	pcData *widget.Label // The label for the value stored in the program counter
	accLabel *widget.Label // The label for the accumulator
	accData *widget.Label // The label for the value stored in the accumulator
	regLabels []*widget.Label // The labels for the registers
	regData []*widget.Label // The labels for the values stored in the registers
	ifidLabels []*widget.Label // The labels for the IFID register
}

// Creates the GuiData struct
func NewGuiData(trackedCpu *cpu.Cpu) *GuiData {
	guiData := GuiData { cpu: trackedCpu}

	// Add the time labels
	guiData.timeLabel = widget.NewLabel("Time")
	guiData.curTime = widget.NewLabel("0")

	// Add the PC labels
	guiData.pcLabel = widget.NewLabel("Program Counter")
	guiData.pcData = widget.NewLabel(util.ConvertToHexUint64(uint64(0), 8))

	// Add the ACC labels
	guiData.accLabel = widget.NewLabel("ACC")
	guiData.accData = widget.NewLabel(util.ConvertToHexUint64(uint64(0), 16))

	// Create the lists of labels for the registers
	guiData.regLabels = make([]*widget.Label, 32)
	guiData.regData = make([]*widget.Label, 32)

	for i := 0; i < len(guiData.regLabels); i++ {
		// Add the special case text for the labels
		specialReg := ""
		switch i {
		case 16:
			specialReg = " (IP0)"
		case 17:
			specialReg = " (IP1)"
		case 28:
			specialReg = " (SP)"
		case 29:
			specialReg = " (FP)"
		case 30:
			specialReg = " (LR)"
		}

		if i == 31 {
			// XZR is not represented by an X31
			guiData.regLabels[i] = widget.NewLabel("XZR")
		} else {
			// Add the label with the special text if needed
			guiData.regLabels[i] = widget.NewLabel(fmt.Sprintf("X%d%s", i, specialReg))
		}

		// Initialize registers to 0
		guiData.regData[i] = widget.NewLabel(util.ConvertToHexUint64(uint64(0), 16))
	}

	// Create the IFID labels
	guiData.ifidLabels = make([]*widget.Label, 2)
	for i := 0; i < len(guiData.ifidLabels); i++ {
		switch i {
		case 0:
			guiData.ifidLabels[i] = widget.NewLabel(fmt.Sprintf("Instruction: %s", util.ConvertToHexUint32(0, 8)))
		case 1:
			guiData.ifidLabels[i] = widget.NewLabel(fmt.Sprintf("Incremented PC: %s", util.ConvertToHexUint32(0, 8)))
		}
	}

	return &guiData
}

// Function that gets called every clock cycle
func (guiData *GuiData) Pulse() {

	// Update register labels with locks
	for i := 0; i < len(guiData.regLabels) - 1; i++ {
		// Add the special case text for the labels
		specialReg := ""
		switch i {
		case 16:
			specialReg = " (IP0)"
		case 17:
			specialReg = " (IP1)"
		case 28:
			specialReg = " (SP)"
		case 29:
			specialReg = " (FP)"
		case 30:
			specialReg = " (LR)"
		}

		lockText := " (Lock)"

		if guiData.cpu.GetRegisterLocks().Contains(uint32(i)) {
			guiData.regLabels[i].SetText(fmt.Sprintf("X%d%s%s", i, specialReg, lockText))
		} else {
			guiData.regLabels[i].SetText(fmt.Sprintf("X%d%s", i, specialReg))
		}
	}

	// Update the register values
	guiData.curTime.SetText(fmt.Sprintf("%d", util.GetCurrentTime()))
	guiData.pcData.SetText(util.ConvertToHexUint32(uint32(guiData.cpu.GetProgramCounter()), 8))
	guiData.accData.SetText(util.ConvertToHexUint64(guiData.cpu.GetAcc(), 16))
	for i := 0; i < len (guiData.regData); i++ {
		guiData.regData[i].SetText(util.ConvertToHexUint64(guiData.cpu.GetRegisters()[i], 16))
	}
	// Update the IFID values
	for i := 0; i < len(guiData.ifidLabels); i++ {
		switch i {
		case 0:
			guiData.ifidLabels[i].SetText(fmt.Sprintf("Instruction: %s", util.ConvertToHexUint32(guiData.cpu.GetIFIDReg().GetInstruction(), 8)))
		case 1:
			guiData.ifidLabels[i].SetText(fmt.Sprintf("Incremented PC: %s", util.ConvertToHexUint32(uint32(guiData.cpu.GetIFIDReg().GetIncrementedPC()), 8)))
		}
	}
}

// Function that initializes and starts the gui
func CreateGui(guiData *GuiData) {
	// Create the app and window
	app := app.New()
	win := app.NewWindow("GoVM")

	// Create the grid and add the labels and data
	grid := container.New(layout.NewGridLayout(4))
	
	grid.Add(guiData.timeLabel)
	grid.Add(guiData.curTime)

	grid.Add(guiData.pcLabel)
	grid.Add(guiData.pcData)

	grid.Add(guiData.accLabel)
	grid.Add(guiData.accData)

	for i := 0; i < len(guiData.regLabels); i++ {
		grid.Add(guiData.regLabels[i])
		grid.Add(guiData.regData[i])
	}

	// Create the IFID grid and accordion item
	ifidTable := container.New(layout.NewGridLayout(1))
	ifidTable.Add(guiData.ifidLabels[0])
	ifidTable.Add(guiData.ifidLabels[1])
	ifidAccordionItem := widget.NewAccordionItem("IFID Register", ifidTable)
	ifidAccordionItem.Open = true

	pipelineAccordion := widget.NewAccordion(ifidAccordionItem)

	content := container.NewHSplit(grid, pipelineAccordion)

	// Add the grid to the window
	win.SetContent(content)

	// Start the application
	win.ShowAndRun()
}