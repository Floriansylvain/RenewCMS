package gateways

import "RenewCMS/domain/user"

type IUserRepository interface {
	Get(id uint32) (user.User, error)
	GetByUsername(username string) (user.User, error)
	GetByEmail(email string) (user.User, error)
	GetAll() []user.User
	Create(user user.User) (user.User, error)
	Delete(id uint32) error
	UpdateVerificationStatus(userId uint32, isVerified bool) (user.User, error)
	UpdatePassword(userId uint32, password string) (user.User, error)
	UpdatePasswordResetCode(userId uint32, code string) (user.User, error)
}
