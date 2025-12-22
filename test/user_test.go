package test

import (
	"RenewCMS/internal/domain/user"
	"RenewCMS/internal/infrastructure/useCases"
	"testing"
)

type mockUserRepo struct {
	users map[uint32]user.User
}

func (m *mockUserRepo) Create(u user.User) (user.User, error) {
	u.ID = uint32(len(m.users) + 1)
	m.users[u.ID] = u
	return u, nil
}

func (m *mockUserRepo) Get(id uint32) (user.User, error) {
	u, ok := m.users[id]
	if !ok {
		return user.User{}, nil
	}
	return u, nil
}

func (m *mockUserRepo) GetAll() []user.User {
	var list []user.User
	for _, u := range m.users {
		list = append(list, u)
	}
	return list
}

func (m *mockUserRepo) UpdateVerificationStatus(id uint32, v bool) (user.User, error) {
	u := m.users[id]
	u.IsVerified = v
	m.users[id] = u
	return u, nil
}

func (m *mockUserRepo) GetByUsername(n string) (user.User, error) { return user.User{}, nil }
func (m *mockUserRepo) GetByEmail(e string) (user.User, error)    { return user.User{}, nil }
func (m *mockUserRepo) Delete(id uint32) error                    { return nil }
func (m *mockUserRepo) UpdatePassword(id uint32, p string) (user.User, error) {
	return user.User{}, nil
}
func (m *mockUserRepo) UpdatePasswordResetCode(id uint32, c string) (user.User, error) {
	return user.User{}, nil
}

var _ user.Repository = &mockUserRepo{}

func TestCreateUser(t *testing.T) {
	repo := &mockUserRepo{users: make(map[uint32]user.User)}
	uc := useCases.NewCreateUserUseCase(repo)

	cmd := useCases.CreateUserCommand{
		Username: "testuser",
		Password: "password123",
		Email:    "test@example.com",
	}

	res, err := uc.CreateUser(cmd)
	if err != nil {
		t.Fatalf("User creation failed: %v", err)
	}
	if res.Username != cmd.Username {
		t.Errorf("Username mismatch")
	}
	if res.VerificationCode == "" {
		t.Errorf("Verification code was not generated")
	}
}

func TestUpdateUserVerification(t *testing.T) {
	repo := &mockUserRepo{users: map[uint32]user.User{
		1: {ID: 1, IsVerified: false},
	}}
	uc := useCases.NewUpdateUserUseCase(repo)

	res, err := uc.UpdateVerificationStatus(1, true)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	if !res.IsVerified {
		t.Errorf("User should be verified")
	}
}
