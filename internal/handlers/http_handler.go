package handlers

import (
	"NanoQR/internal/model"
	"NanoQR/internal/service"
	"net/http"
	"encoding/json"
	"encoding/base64"
    "strings"
)

var allowedOrigins = []string {
	"https://nanoqr-web.vercel.app",
	"http://localhost:5173",
	"https://another-web.com",
}

func isOriginAllowed(origin string) bool {

	for _, o := range allowedOrigins {
		if o == origin {
			return true
		}
	}
	return false

}

func CORSMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		origin := r.Header.Get("Origin")
		if isOriginAllowed(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Content-Type", "application/json")
		}

		// Preflight
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }

		// Call next handler (QRhandler)
		next.ServeHTTP(w, r)
	})

}

// Create qrService variable for QRService interface with DefualtQRService implementation
var qrService service.QRService = &service.DefaultQRService{}

func QRhandler (w http.ResponseWriter, r *http.Request) {
	
	

	// Define entryData as the JsonRequest struct in model/model.go
	var entryData model.JsonRequest

	// If the method isn't POST
	if r.Method != http.MethodPost {
		http.Error(w, "Not a valid method. Use POST", http.StatusMethodNotAllowed)
		return
	}

	// Decode the json file and forward the data to entryData (model.JsonRequest).
	// With "err :=" we are checking if it match with the struct model.
	err := json.NewDecoder(r.Body).Decode(&entryData)

		// If it doesn't match:
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// (local handler error) If input is empty: 
		if strings.TrimSpace(entryData.Input) == "" {
		http.Error(w, "Input cannot be empty", http.StatusBadRequest)
		return
		}

		// If math, create the QR:
		qr, finalSize, finalRecoverLevel, err := qrService.Generate(entryData.Input, entryData.Size, entryData.RecoverLevel)

			// Error generating QR:
			if err != nil {
				http.Error(w, "Error generating QR", http.StatusBadRequest)
				return
			}

			// Succes: we encode the QR bytes to Base64:
			qrBase64 := base64.StdEncoding.EncodeToString(qr)

			// Then build the response for the client:
			response := map[string]any{
				"status": "success",
				"info": map[string]any{
					"input": entryData.Input,
					"size": finalSize,
					"recoverLevel": finalRecoverLevel,

				},
				"qr": qrBase64,
			}

			// Here we encode the response in JSON format and send it to the client:
			json.NewEncoder(w).Encode(response)
}