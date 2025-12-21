package useCases

import "RenewCMS/internal/domain/user"

type UpdateUserUseCase struct {
	userRepository user.Repository
}

func NewUpdateUserUseCase(userRepository user.Repository) *UpdateUserUseCase {
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
