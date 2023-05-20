package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"time"
)

func ReadFileToBuffer(filename string) (*bytes.Buffer, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(fileBytes)
	return buffer, nil
}

func AppendBufferToFile(filename string, buffer *bytes.Buffer) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = buffer.WriteTo(file)
	if err != nil {
		return err
	}

	return nil
}

func CreateFileName() string {
	currentTime := time.Now()
	filename := fmt.Sprintf("file_%s", currentTime.Format("20060102_150405"))
	return filename
}

func FileToBase64(filename string) (string, error) {
	fileBytes, err := os.ReadFile(filename)
	println(filename)
	if err != nil {
		println("Couldn't read the file")
		return "", err
	}

	encodedString := base64.StdEncoding.EncodeToString(fileBytes)
	return encodedString, nil
}
