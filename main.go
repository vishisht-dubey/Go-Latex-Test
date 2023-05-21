package main

import (
	section "Go-Latex-Test/section"
	utils "Go-Latex-Test/utils"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/json-to-latex", getLatex)

	router.Run("localhost:3000")
}

func getLatex(c *gin.Context) {
	var resume section.Resume

	if err := c.BindJSON(&resume); err != nil {
		fmt.Println(err)
		return
	}

	outputDir := "/tmp/"
	jobName := utils.CreateFileName()
	texFile := outputDir + jobName + ".tex"
	pdfFile := outputDir + jobName + ".pdf"

	resumeBuffer, _ := resume.Prepare()
	utils.AppendBufferToFile(texFile, resumeBuffer)

	if stdout, err := exec.Command("pdflatex", "-output-directory="+"/tmp/", "-jobname="+jobName, texFile).Output(); err != nil {
		println(err.Error())
	} else {
		println(string(stdout))
	}

	b64, _ := utils.FileToBase64(pdfFile)

	c.IndentedJSON(http.StatusCreated, gin.H{"data": b64})
}
