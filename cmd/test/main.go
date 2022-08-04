package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joshuaseligman/GoVM/pkg/assembler"
	"github.com/joshuaseligman/GoVM/pkg/hardware/memory"
)

type ProgStruct struct {
	Prog string `json:"prog"`
}

type StatusStruct struct {
	Memory *memory.MemoryAPI `json:"memory"`
}

func main() {
	r := gin.Default()

	r.POST("/api/asmprog", handleAsmProg)
	r.GET("/api/test", test)

	r.Use(cors.Default())
	r.Run(":8080")
}

func handleAsmProg(c *gin.Context) {
	// Get the data from the request body
	if raw, err := c.GetRawData(); err != nil {
		// Return error if there is an error
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusBadRequest, gin.H {
			"err": err.Error(),
		})
	} else {
		// Convert the raw data into the struct for the program
		var prog ProgStruct
		json.Unmarshal(raw, &prog)

		// Assemble the program
		if binaryProg, err := assembler.AssembleProgramAPI(prog.Prog); err == nil {
			// Return the binary
			c.Header("Access-Control-Allow-Origin", "*")
			c.JSON(http.StatusOK, gin.H {
				"binaryProg": binaryProg,
			})
		} else {
			// Return the error
			c.Header("Access-Control-Allow-Origin", "*")
			c.JSON(http.StatusBadRequest, gin.H {
				"err": err.Error(),
			})
		}
	}
}

func test(c *gin.Context) {
	mem := memory.NewEmptyMemory(10)

	// Update users every 1 second
	ticker := time.NewTicker(1 * time.Second)
	defer func() {
		ticker.Stop()
	}()

	// Stream the updated info to the user
	c.Stream(func(w io.Writer) bool {
		select {
		case <-ticker.C:
			newStatus := StatusStruct {
				Memory: mem.ConvertAPI(),
			}
			// Send the "ping" event to the user with the updated queue
			c.Header("Access-Control-Allow-Origin", "*")
			c.SSEvent("ping", newStatus)
		}
		return true
	})
}