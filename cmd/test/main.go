package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joshuaseligman/GoVM/pkg/assembler"
)

type ProgStruct struct {
	Prog string `json:"prog"`
}

func main() {
	r := gin.Default()

	r.POST("/api/asmprog", handleAsmProg)

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