package service

import (
	qrcode "github.com/skip2/go-qrcode"

)

// Here we're gonna create the struct that  	
type DefaultQRService struct{}

func (s *DefaultQRService) Generate(input string, size int, recoverLevel string) ([]byte, error) {

	qrRecoverLevel := qrcode.Medium

	// QR Recover levels:
	switch recoverLevel {
	case "low":
		qrRecoverLevel = qrcode.Low
	case "medium":
		qrRecoverLevel = qrcode.Medium
	case "high":
		qrRecoverLevel = qrcode.High
	case "highest":
		qrRecoverLevel = qrcode.Highest
	}

	// Size limits & default:
	if size <= 0 {
		size = 256
	}
	if size < 41 {
		size = 41
	}
	if size > 2048 {
		size = 2048
	}

	return qrcode.Encode(input, qrRecoverLevel, size)

}