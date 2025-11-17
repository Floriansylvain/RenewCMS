package useCases

import (
	"RenewCMS/domain/gateways"
)

type SendMailUseCase struct {
	mailRepository gateways.IMailRepository
}

func NewSendMailUseCase(mailRepository gateways.IMailRepository) *SendMailUseCase {
	return &SendMailUseCase{mailRepository}
}

func (g *SendMailUseCase) SendMail(receiverAddress string, templateName string, data any) error {
	return g.mailRepository.Send(receiverAddress, templateName, data)
}
