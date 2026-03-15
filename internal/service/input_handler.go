package service

import (
	"bufio"
	"io"
	"fmt"
)

// String reader
func ReadString(r io.Reader, message string) (string, error) {

	scanner := bufio.NewScanner(r)
	fmt.Print(message)

	if scanner.Scan() {
		return scanner.Text(), nil
	}

	// Default (scanner error)
	return "", fmt.Errorf("Failed to read string")
}
