package useCases

import (
	"RenewCMS/internal/domain/user"
	"time"

	"github.com/google/uuid"
)

type CreateUserUseCase struct {
	userRepository user.Repository
}

type CreateUserCommand struct {
	Username string
	Password string
	Email    string
}

func NewCreateUserUseCase(userRepository user.Repository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository}
}

func (g *CreateUserUseCase) CreateUser(createUser CreateUserCommand) (user.User, error) {
	rawUuid := uuid.NewString()

	createdUser, err := g.userRepository.Create(user.User{
		Username:               createUser.Username,
		Password:               createUser.Password,
		PasswordResetCode:      "",
		Email:                  createUser.Email,
		IsVerified:             false,
		VerificationCode:       rawUuid,
		VerificationExpiration: time.Now().Add(2 * time.Hour),
	})
	if err != nil {
		return user.User{}, err
	}

	createdUser.VerificationCode = rawUuid
	return createdUser, nil
}
