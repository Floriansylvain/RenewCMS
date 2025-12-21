package mappers

import (
	"RenewCMS/internal/domain/user"
	"RenewCMS/internal/infrastructure/persistence/models"
)

func UserToDomain(m models.User) user.User {
	return user.User{
		ID:                     m.ID,
		Username:               m.Username,
		Password:               m.Password,
		PasswordResetCode:      m.PasswordResetCode,
		Email:                  m.Email,
		IsVerified:             m.IsVerified,
		VerificationCode:       m.VerificationCode,
		VerificationExpiration: m.VerificationExpiration,
		CreatedAt:              m.CreatedAt,
		UpdatedAt:              m.UpdatedAt,
	}
}

func UserToModel(a user.User) models.User {
	return models.User{
		Username:               a.Username,
		Password:               a.Password,
		PasswordResetCode:      a.PasswordResetCode,
		Email:                  a.Email,
		IsVerified:             a.IsVerified,
		VerificationCode:       a.VerificationCode,
		VerificationExpiration: a.VerificationExpiration,
		CreatedAt:              a.CreatedAt,
		UpdatedAt:              a.UpdatedAt,
	}
}
