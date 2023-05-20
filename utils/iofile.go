package utils

import (
	"bytes"
	"os"
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
