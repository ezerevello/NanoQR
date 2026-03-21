package service

type QRService interface {
	Generate(input string, size int, recoverLevel string) ([]byte, int, string, error)
}