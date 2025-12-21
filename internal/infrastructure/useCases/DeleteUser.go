package useCases

import "RenewCMS/internal/domain/user"

type DeleteUserUseCase struct {
	userRepository user.Repository
}

func NewDeleteUserUseCase(userRepository user.Repository) *DeleteUserUseCase {
	return &DeleteUserUseCase{userRepository}
}

func (g *DeleteUserUseCase) DeleteUser(userId uint32) error {
	return g.userRepository.Delete(userId)
}
