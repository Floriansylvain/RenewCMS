package useCases

import (
	"RenewCMS/domain/gateways"
	"RenewCMS/domain/user"
)

type UpdateUserUseCase struct {
	userRepository gateways.IUserRepository
}

func NewUpdateUserUseCase(userRepository gateways.IUserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{userRepository}
}

func (g *UpdateUserUseCase) UpdateVerificationStatus(userId uint32, isVerified bool) (user.User, error) {
	return g.userRepository.UpdateVerificationStatus(userId, isVerified)
}

func (g *UpdateUserUseCase) UpdatePasswordResetCode(userId uint32, code string) (user.User, error) {
	return g.userRepository.UpdatePasswordResetCode(userId, code)
}

func (g *UpdateUserUseCase) UpdatePassword(userId uint32, password string) (user.User, error) {
	return g.userRepository.UpdatePassword(userId, password)
}
