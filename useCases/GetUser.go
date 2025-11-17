package useCases

import (
	"RenewCMS/domain/gateways"
	"RenewCMS/domain/user"
)

type GetUserUseCase struct {
	userRepository gateways.IUserRepository
}

func NewGetUserUseCase(userRepository gateways.IUserRepository) *GetUserUseCase {
	return &GetUserUseCase{userRepository}
}

func (g *GetUserUseCase) GetUser(id uint32) (user.User, error) {
	return g.userRepository.Get(id)
}

func (g *GetUserUseCase) GetUserByUsername(username string) (user.User, error) {
	return g.userRepository.GetByUsername(username)
}

func (g *GetUserUseCase) GetUserByEmail(email string) (user.User, error) {
	return g.userRepository.GetByEmail(email)
}
