package service

import (
	qrcode "github.com/skip2/go-qrcode"

)

// Here we're gonna create the struct that  	
type DefaultQRService struct{}

func (s *DefaultQRService) Generate(input string, size int, recoverLevel string) ([]byte, int, string, error) {
	
    qrRecoverLevel := qrcode.Medium
    finalRecoverLevel := recoverLevel

    // Defaults for recoverLevel
    if recoverLevel == "" {
        finalRecoverLevel = "medium"
    }

    switch finalRecoverLevel {
    case "low":
        qrRecoverLevel = qrcode.Low
    case "medium":
        qrRecoverLevel = qrcode.Medium
    case "high":
        qrRecoverLevel = qrcode.High
    case "highest":
        qrRecoverLevel = qrcode.Highest
    default:
        finalRecoverLevel = "medium"
        qrRecoverLevel = qrcode.Medium
    }

    // Size limits & default:
    finalSize := size
    if finalSize <= 0 {
        finalSize = 256
    }
    if finalSize < 41 {
        finalSize = 41
    }
    if finalSize > 2048 {
        finalSize = 2048
    }

    qr, err := qrcode.Encode(input, qrRecoverLevel, finalSize)
    return qr, finalSize, finalRecoverLevel, err

}