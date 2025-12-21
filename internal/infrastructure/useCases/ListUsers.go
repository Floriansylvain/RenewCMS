package useCases

import "RenewCMS/internal/domain/user"

type ListUsersUseCase struct {
	userRepository user.Repository
}

func NewListUsersUseCase(userRepository user.Repository) *ListUsersUseCase {
	return &ListUsersUseCase{userRepository}
}

func (g *ListUsersUseCase) ListUsers() []user.User {
	return g.userRepository.GetAll()
}
