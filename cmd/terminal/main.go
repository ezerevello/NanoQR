package main

import (
	"NanoQR/internal/service"
	"fmt"
	"os"
	qrcode "github.com/skip2/go-qrcode"
)

func main () {
	// Asks for the URL
	input, err := service.ReadString(os.Stdin, "Enter a link: ")

	// Error
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	
	// Succes
	qr, err := qrcode.Encode(input, qrcode.Medium, 256)

		// Error generating QR code
		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		// Succes
		os.WriteFile("test.png", qr, 0644)
		fmt.Println("QR Generated succesfully O_o")
}