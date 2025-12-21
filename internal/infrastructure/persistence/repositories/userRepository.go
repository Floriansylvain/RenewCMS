package persistence

import (
	domainUser "RenewCMS/internal/domain/user"
	"RenewCMS/internal/infrastructure/persistence/mappers"
	entity "RenewCMS/internal/infrastructure/persistence/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) Get(id uint32) (domainUser.User, error) {
	var localUser entity.User
	err := u.db.Model(&entity.User{}).First(&localUser, id).Error
	if err != nil {
		return domainUser.User{}, err
	}

	return mappers.UserToDomain(localUser), nil
}

func (u *UserRepository) Create(user domainUser.User) (domainUser.User, error) {
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
		return domainUser.User{}, creationResult.Error
	}

	var createdUser entity.User
	creationResult.Scan(&createdUser)

	return mappers.UserToDomain(createdUser),
		nil
}

func (u *UserRepository) Delete(id uint32) error {
	return u.db.Delete(&domainUser.User{}, id).Error
}

func (u *UserRepository) GetAll() []domainUser.User {
	var users []entity.User
	u.db.Model(&entity.User{}).Find(&users)

	var domainUsers []domainUser.User
	for _, localUser := range users {
		domainUsers = append(domainUsers, mappers.UserToDomain(localUser))
	}

	return domainUsers
}

func (u *UserRepository) GetByUsername(username string) (domainUser.User, error) {
	var localUser entity.User
	err := u.db.Model(&entity.User{}).Where("username = ?", username).First(&localUser).Error
	if err != nil {
		return domainUser.User{}, err
	}

	return mappers.UserToDomain(localUser), nil
}

func (u *UserRepository) GetByEmail(email string) (domainUser.User, error) {
	var localUser entity.User
	err := u.db.Model(&entity.User{}).Where("email = ?", email).First(&localUser).Error
	if err != nil {
		return domainUser.User{}, err
	}

	return mappers.UserToDomain(localUser), nil
}

func (u *UserRepository) UpdateVerificationStatus(userId uint32, isVerified bool) (domainUser.User, error) {
	var localUser entity.User
	err := u.db.Model(&entity.User{}).First(&localUser, userId).Error
	if err != nil {
		return domainUser.User{}, err
	}

	localUser.IsVerified = isVerified
	err = u.db.Save(&localUser).Error
	if err != nil {
		return domainUser.User{}, err
	}

	return mappers.UserToDomain(localUser), nil
}

func (u *UserRepository) UpdatePassword(userId uint32, password string) (domainUser.User, error) {
	var localUser entity.User
	err := u.db.Model(&entity.User{}).First(&localUser, userId).Error
	if err != nil {
		return domainUser.User{}, err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	localUser.Password = string(hashedPassword)
	err = u.db.Save(&localUser).Error
	if err != nil {
		return domainUser.User{}, err
	}

	return mappers.UserToDomain(localUser), nil
}

func (u *UserRepository) UpdatePasswordResetCode(userId uint32, code string) (domainUser.User, error) {
	var localUser entity.User
	err := u.db.Model(&entity.User{}).First(&localUser, userId).Error
	if err != nil {
		return domainUser.User{}, err
	}

	hashedCode, _ := bcrypt.GenerateFromPassword([]byte(code), 12)
	localUser.PasswordResetCode = string(hashedCode)
	err = u.db.Save(&localUser).Error
	if err != nil {
		return domainUser.User{}, err
	}

	return mappers.UserToDomain(localUser), nil
}

var _ domainUser.Repository = &UserRepository{}
