package components

import "fmt"

type AbstractMailService interface {
	SendMail()
}

type AzureMailService struct {
}

func NewAzureMailService() *AzureMailService {
	return &AzureMailService{}
}

func (ams *AzureMailService) SendMail() {
	fmt.Println("Azure sending mail")
}

type GcpMailService struct {
}

func NewGcpMailService() *GcpMailService {
	return &GcpMailService{}
}

func (ams *GcpMailService) SendMail() {
	fmt.Println("Google sending mail")
}
