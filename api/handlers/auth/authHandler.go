package auth

import (
	"RenewCMS/api/pkg/apperr"
	validator "RenewCMS/api/pkg/validator"
	"RenewCMS/internal/domain/user"
	"RenewCMS/internal/infrastructure/useCases"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	Token         *jwtauth.JWTAuth
	GetUseCase    *useCases.GetUserUseCase
	ListUseCase   *useCases.ListUsersUseCase
	CreateUseCase *useCases.CreateUserUseCase
}

type RegisterCredentials struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
}

type LoginCredentials struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8"`
}

func shouldCookieBeSecure() bool {
	return os.Getenv("ENVIRONMENT") == "production"
}

func removeJwtCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now(),
		MaxAge:   -1,
		Secure:   shouldCookieBeSecure(),
		HttpOnly: true,
		Path:     "/",
	})
}

func NewHandler(token *jwtauth.JWTAuth, get *useCases.GetUserUseCase, list *useCases.ListUsersUseCase, create *useCases.CreateUserUseCase) *Handler {
	return &Handler{
		Token:         token,
		GetUseCase:    get,
		ListUseCase:   list,
		CreateUseCase: create,
	}
}

func (h *Handler) SetJwtCookie(w *http.ResponseWriter, userId uint32) error {
	_, tokenString, err := h.Token.Encode(map[string]any{"user_id": userId})
	if err != nil {
		return err
	}
	http.SetCookie(*w, &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   shouldCookieBeSecure(),
		HttpOnly: true,
		Path:     "/",
	})
	return nil
}

func (h *Handler) SomeUsersVerified() bool {
	users := h.ListUseCase.ListUsers()
	for _, localUser := range users {
		if localUser.IsVerified {
			return true
		}
	}
	return false
}

func (h *Handler) IsUserTableEmpty() bool {
	users := h.ListUseCase.ListUsers()
	return len(users) == 0
}

func (h *Handler) IsLoggedIn(r *http.Request) bool {
	token, err := jwtauth.VerifyRequest(
		h.Token,
		r,
		jwtauth.TokenFromCookie,
		jwtauth.TokenFromHeader,
		jwtauth.TokenFromQuery)
	return token != nil && err == nil
}

func (h *Handler) IsVerified(r *http.Request) bool {
	token, err := jwtauth.VerifyRequest(
		h.Token,
		r,
		jwtauth.TokenFromCookie,
		jwtauth.TokenFromHeader,
		jwtauth.TokenFromQuery)
	if err != nil {
		return false
	}
	userId := token.PrivateClaims()["user_id"].(float64)
	currentUser, _ := h.GetUseCase.GetUser(uint32(userId))
	return currentUser.IsVerified
}

func (h *Handler) GetUserFromCredentials(credentials LoginCredentials) (user.User, error) {
	dbUser, err := h.GetUseCase.GetUserByUsername(credentials.Username)
	if err != nil {
		return user.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(credentials.Password))
	if err != nil {
		return user.User{}, err
	}

	return dbUser, nil
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials LoginCredentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, apperr.BodyErrorMessage, http.StatusBadRequest)
		return
	}

	err = validator.Validate.Struct(credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbUser, err := h.GetUserFromCredentials(credentials)
	if err != nil {
		http.Error(w, apperr.LogsErrorMessage, http.StatusForbidden)
		return
	}

	_ = h.SetJwtCookie(&w, dbUser.ID)

	message, _ := json.Marshal(map[string]any{"message": "User logged in! HTTPonly jwt cookie created"})
	_, _ = w.Write(message)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if h.SomeUsersVerified() {
		http.Error(w, "You are not allowed to create a user. Log in or reset database.", http.StatusForbidden)
		return
	}

	var credentials RegisterCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, apperr.BodyErrorMessage, http.StatusBadRequest)
		return
	}

	err = validator.Validate.Struct(credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := h.CreateUseCase.CreateUser(useCases.CreateUserCommand{
		Username: credentials.Username,
		Password: credentials.Password,
		Email:    credentials.Email,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = h.SetJwtCookie(&w, createdUser.ID)

	message, _ := json.Marshal(map[string]any{"message": "User registered! HTTPonly jwt cookie created"})
	_, _ = w.Write(message)
}

func (h *Handler) Logout(w http.ResponseWriter, _ *http.Request) {
	removeJwtCookie(w)
	message, _ := json.Marshal(map[string]any{"message": "User logged out! HTTPonly jwt cookie deleted"})
	_, _ = w.Write(message)
}
