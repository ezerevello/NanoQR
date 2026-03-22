package main

import (
	"net/http"
	"fmt"
	"log"
	"NanoQR/internal/handlers"
)

func main () {
	// Any request for "api/qr" goes to service.QRhandler:
	http.HandleFunc("/api/qr", handlers.QRhandler)

	// This endpoint is for test the API
    http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK\n"))
    })

	// Port config:
	port := ":8080"
	fmt.Printf("Starting NanoQR in http://localhost%s\n", port)

	// This prevent the app from closing because it's listening 24/7 to the port:
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("The server crashed: %v", err)
	}
}