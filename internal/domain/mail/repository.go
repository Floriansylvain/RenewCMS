package mail

type Repository interface {
	Send(receiverAddress string, templateName string, data any) error
}
