package utils

import (
	"bufio"
	"io"
	"os"
)

// Reads all contents of the file using bufio reader
func ReadFromFile(file *os.File) ([]byte, error) {

	reader := bufio.NewReader(file)
	data := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		data = append(data, buf[:n]...)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
	}

	return data, nil
}

// returns TRUE if file exists. False if not.
func CheckFileExistence(path string) bool {

	_, pathErr := os.Stat(path)
	if pathErr != nil {
		if os.IsNotExist(pathErr) {
			return false
		}
	}

	return true
}
