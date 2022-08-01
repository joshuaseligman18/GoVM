package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshuaseligman/GoVM/pkg/assembler"
)

func main() {
	r := gin.Default()

	r.POST("/api/asmprog", handleAsmProg)

	r.Run(":8080")
}

func handleAsmProg(c *gin.Context) {
	c.Request.ParseForm()
	program := c.Request.PostForm.Get("prog")
	bin, err := assembler.AssembleProgramAPI(program)
	if err == nil {
		c.JSON(http.StatusOK, gin.H {
			"progBinary": bin,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H {
			"msg": err.Error(),
		})
	}
}