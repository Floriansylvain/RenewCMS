package user

type Repository interface {
	Get(id uint32) (User, error)
	GetByUsername(username string) (User, error)
	GetByEmail(email string) (User, error)
	GetAll() []User
	Create(user User) (User, error)
	Delete(id uint32) error
	UpdateVerificationStatus(userId uint32, isVerified bool) (User, error)
	UpdatePassword(userId uint32, password string) (User, error)
	UpdatePasswordResetCode(userId uint32, code string) (User, error)
}
