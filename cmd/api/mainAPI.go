package main

import (
	// "github.com/joshuaseligman/GoVM/pkg/assembler"
	// "github.com/joshuaseligman/GoVM/pkg/hardware/clock"
	// "github.com/joshuaseligman/GoVM/pkg/hardware/cpu"
	// "github.com/joshuaseligman/GoVM/pkg/hardware/memory"

	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// assembledProgram := assembler.AssembleProgram("test.goas", 0x10000)
	
	// mem := memory.NewFlashedMemory(assembledProgram)
	
	// mem.MemoryDump(0, 30)
	
	// cpu := cpu.NewCpu(mem)

	// clk := clock.NewClock()

	// clk.AddClockListener(cpu)
	
	// go clk.StartClock(500)

	r := gin.Default()

	r.GET("/", test)

	r.Run(":8080")
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"msg": "hello",
	})
}