package useCases

import "RenewCMS/internal/domain/mail"

type SendMailUseCase struct {
	mailRepository mail.Repository
}

func NewSendMailUseCase(mailRepository mail.Repository) *SendMailUseCase {
	return &SendMailUseCase{mailRepository}
}

func (g *SendMailUseCase) SendMail(receiverAddress string, templateName string, data any) error {
	return g.mailRepository.Send(receiverAddress, templateName, data)
}
