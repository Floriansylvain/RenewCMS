package test

import (
	"RenewCMS/internal/domain/mail"
	"RenewCMS/internal/infrastructure/useCases"
	"errors"
	"testing"
)

type mockMailRepo struct {
	SentEmails []struct {
		Address  string
		Template string
		Data     any
	}
	ShouldFail bool
}

func (m *mockMailRepo) Send(receiverAddress string, templateName string, data any) error {
	if m.ShouldFail {
		return errors.New("smtp error")
	}
	m.SentEmails = append(m.SentEmails, struct {
		Address  string
		Template string
		Data     any
	}{receiverAddress, templateName, data})
	return nil
}

var _ mail.Repository = &mockMailRepo{}

func TestSendMail(t *testing.T) {
	repo := &mockMailRepo{}
	uc := useCases.NewSendMailUseCase(repo)

	t.Run("Successful Email Send", func(t *testing.T) {
		repo.ShouldFail = false
		err := uc.SendMail("user@example.com", "welcome", "User123")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if len(repo.SentEmails) != 1 {
			t.Errorf("Expected 1 email to be sent, got %d", len(repo.SentEmails))
		}
	})

	t.Run("Failed Email Send", func(t *testing.T) {
		repo.ShouldFail = true
		err := uc.SendMail("fail@example.com", "reset", nil)
		if err == nil {
			t.Errorf("Expected an error but got nil")
		}
	})
}
