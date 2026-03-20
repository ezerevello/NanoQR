package service

import (
	"NanoQR/internal/model"
	"net/http"
	"encoding/json"
	"encoding/base64"
	qrcode "github.com/skip2/go-qrcode"
)

func QRhandler (w http.ResponseWriter, r *http.Request) {

	// If the method isn't POST
	if r.Method != http.MethodPost {
		http.Error(w, "Not a valid method. Use POST", http.StatusMethodNotAllowed)
		return
	}

	// Define entryData as the JsonRequest struct in model/model.go
	var entryData model.JsonRequest

	// Decode the json file and forward the data to entryData (model.JsonRequest).
	// With "err :=" we are checking if it match with the struct model.
	err := json.NewDecoder(r.Body).Decode(&entryData)

		// If it doesn't match:
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// If match, first we declare the entryData.RecoverLevel default value:
		qrRecoverLevel := qrcode.Medium

		// Then specify the levels:
		switch entryData.RecoverLevel {
		case "low":
			qrRecoverLevel = qrcode.Low
		case "medium":
			qrRecoverLevel = qrcode.Medium
		case "high":
			qrRecoverLevel = qrcode.High
		case "highest":
			qrRecoverLevel = qrcode.Highest
		}

		// Then define the size limits & default:
		if entryData.Size <= 0 {
			entryData.Size = 256
		}
		if entryData.Size < 41 {
			entryData.Size = 41
		}
		if entryData.Size > 2048 {
			entryData.Size = 2048
		}


		// Now we can create the QR without worry.
		qr, err := qrcode.Encode(entryData.Input, qrRecoverLevel, entryData.Size)

			// Error generating QR:
			if err != nil {
				http.Error(w, "Error generating QR", http.StatusBadRequest)
				return
			}

			// Succes:
			w.Header().Set("Content-Type", "application/json")

	// After success, we encode the QR bytes to Base64:
	qrBase64 := base64.StdEncoding.EncodeToString(qr)

	// Then build the response for the client:
	response := map[string]any{
		"status": "success",
		"info": map[string]any{
			"input": entryData.Input,
			"size": entryData.Size,
			"recoveryLevel": entryData.RecoverLevel,

		},
		"qr": qrBase64,
	}

	// Here we encode the response in JSON format and send it to the client:
	json.NewEncoder(w).Encode(response)
}