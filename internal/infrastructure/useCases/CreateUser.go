package useCases

import (
	"RenewCMS/internal/domain/user"

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

	createdUser, err := g.userRepository.Create(user.FromApi(
		createUser.Username,
		createUser.Password,
		"",
		createUser.Email,
		rawUuid,
	))
	if err != nil {
		return user.User{}, err
	}

	createdUser.VerificationCode = rawUuid
	return createdUser, nil
}
