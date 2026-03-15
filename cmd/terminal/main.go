package main

import (
	"NanoQR/internal/service"
	"fmt"
	"os"
	qrcode "github.com/skip2/go-qrcode"
)

func main () {
	// Asks for the URL
	input, err := service.ReadString(os.Stdin, "Ingresa el enlace: ")

	// Error
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	
	// Succes
	png, err := qrcode.Encode(input, qrcode.Medium, 256)

		// Error generating QR code
		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		// Succes
		fmt.Print("qr:\n", png)
}