package main

import (
	section "Go-Latex-Test/section"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resume struct {
	SE []section.Education `json:"education"`
}

func main() {
	router := gin.Default()
	router.POST("/get-latex", getLatex)

	router.Run("localhost:3000")
}

func getLatex(c *gin.Context) {
	var resume Resume
	fmt.Println("something is going on")

	if err := c.BindJSON(&resume); err != nil {
		fmt.Println(err)
		return
	}

	// Add the new album to the slice.
	educationSection, _ := section.PrepareEducationSection(resume.SE[0])
	c.IndentedJSON(http.StatusCreated, gin.H{"latex": educationSection})
}
