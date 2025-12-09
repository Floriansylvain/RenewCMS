package useCases

import (
	"RenewCMS/domain/gateways"
	"RenewCMS/domain/user"

	"github.com/google/uuid"
)

type CreateUserUseCase struct {
	userRepository gateways.IUserRepository
}

type CreateUserCommand struct {
	Username string
	Password string
	Email    string
}

func NewCreateUserUseCase(userRepository gateways.IUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository}
}

func (g *CreateUserUseCase) CreateUser(createUser CreateUserCommand) (user.User, error) {
	return g.userRepository.Create(user.FromApi(
		createUser.Username,
		createUser.Password,
		"",
		createUser.Email,
		uuid.NewString(),
	))
}
