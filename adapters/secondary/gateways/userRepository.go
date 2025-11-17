package gateways

import (
	entity "RenewCMS/adapters/secondary/gateways/models"
	"RenewCMS/domain/gateways"
	"RenewCMS/domain/user"
	domain "RenewCMS/domain/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func mapUserToDomain(user entity.User) domain.User {
	return domain.FromDb(
		user.ID,
		user.Username,
		user.Password,
		user.PasswordResetCode,
		user.Email,
		user.IsVerified,
		user.VerificationCode,
		user.VerificationExpiration,
		user.CreatedAt, user.UpdatedAt,
	)
}

func (u *UserRepository) Get(id uint32) (domain.User, error) {
	var localUser entity.User
	err := u.db.Model(&entity.User{}).First(&localUser, id).Error
	if err != nil {
		return domain.User{}, err
	}

	return mapUserToDomain(localUser), nil
}

func (u *UserRepository) Create(user domain.User) (domain.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	hashedVerificationCode, _ := bcrypt.GenerateFromPassword([]byte(user.VerificationCode), 12)

	creationResult := u.db.Create(&entity.User{
		Username:               user.Username,
		Password:               string(hashedPassword),
		Email:                  user.Email,
		VerificationCode:       string(hashedVerificationCode),
		VerificationExpiration: user.VerificationExpiration,
	})
	if creationResult.Error != nil {
		return domain.User{}, creationResult.Error
	}

	var createdUser entity.User
	creationResult.Scan(&createdUser)

	return mapUserToDomain(createdUser),
		nil
}

func (u *UserRepository) Delete(id uint32) error {
	return u.db.Delete(&user.User{}, id).Error
}

func (u *UserRepository) GetAll() []domain.User {
	var users []entity.User
	u.db.Model(&entity.User{}).Find(&users)

	var domainUsers []domain.User
	for _, localUser := range users {
		domainUsers = append(domainUsers, mapUserToDomain(localUser))
	}

	return domainUsers
}

func (u *UserRepository) GetByUsername(username string) (domain.User, error) {
	var localUser entity.User
	err := u.db.Model(&entity.User{}).Where("username = ?", username).First(&localUser).Error
	if err != nil {
		return domain.User{}, err
	}

	return mapUserToDomain(localUser), nil
}

func (u *UserRepository) GetByEmail(email string) (domain.User, error) {
	var localUser entity.User
	err := u.db.Model(&entity.User{}).Where("email = ?", email).First(&localUser).Error
	if err != nil {
		return domain.User{}, err
	}

	return mapUserToDomain(localUser), nil
}

func (u *UserRepository) UpdateVerificationStatus(userId uint32, isVerified bool) (domain.User, error) {
	var localUser entity.User
	err := u.db.Model(&entity.User{}).First(&localUser, userId).Error
	if err != nil {
		return domain.User{}, err
	}

	localUser.IsVerified = isVerified
	err = u.db.Save(&localUser).Error
	if err != nil {
		return domain.User{}, err
	}

	return mapUserToDomain(localUser), nil
}

func (u *UserRepository) UpdatePassword(userId uint32, password string) (domain.User, error) {
	var localUser entity.User
	err := u.db.Model(&entity.User{}).First(&localUser, userId).Error
	if err != nil {
		return domain.User{}, err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	localUser.Password = string(hashedPassword)
	err = u.db.Save(&localUser).Error
	if err != nil {
		return domain.User{}, err
	}

	return mapUserToDomain(localUser), nil
}

func (u *UserRepository) UpdatePasswordResetCode(userId uint32, code string) (domain.User, error) {
	var localUser entity.User
	err := u.db.Model(&entity.User{}).First(&localUser, userId).Error
	if err != nil {
		return domain.User{}, err
	}

	hashedCode, _ := bcrypt.GenerateFromPassword([]byte(code), 12)
	localUser.PasswordResetCode = string(hashedCode)
	err = u.db.Save(&localUser).Error
	if err != nil {
		return domain.User{}, err
	}

	return mapUserToDomain(localUser), nil
}

var _ gateways.IUserRepository = &UserRepository{}
