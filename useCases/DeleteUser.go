package useCases

import (
	"RenewCMS/domain/gateways"
)

type DeleteUserUseCase struct {
	userRepository gateways.IUserRepository
}

func NewDeleteUserUseCase(userRepository gateways.IUserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{userRepository}
}

func (g *DeleteUserUseCase) DeleteUser(userId uint32) error {
	return g.userRepository.Delete(userId)
}
