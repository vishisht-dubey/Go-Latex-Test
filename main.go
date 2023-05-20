package main

import (
	section "Go-Latex-Test/section"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/json-to-latex", getLatex)

	router.Run("localhost:3000")
}

func getLatex(c *gin.Context) {
	var resume section.Resume
	fmt.Println("something is going on")

	if err := c.BindJSON(&resume); err != nil {
		fmt.Println(err)
		return
	}

	// Add the new album to the slice.
	educationSection, _ := section.PrepareEducationSection(resume.Education)
	c.IndentedJSON(http.StatusCreated, gin.H{"latex": educationSection})
}
