package auth

import (
	"RenewCMS/api"
	"RenewCMS/domain/user"
	"RenewCMS/useCases"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterCredentials struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
}

type LoginCredentials struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8"`
}

var shouldCookieBeSecure = os.Getenv("ENVIRONMENT") == "production"
var Token *jwtauth.JWTAuth

const LogsErrorMessage = "Access to the requested resource is forbidden due to incorrect password and/or username."
const BodyErrorMessage = "The request cannot be processed due to a mismatch in the format of the body."

func SetJwtCookie(w *http.ResponseWriter, userId uint32) error {
	_, tokenString, err := Token.Encode(map[string]any{"user_id": userId})
	if err != nil {
		return err
	}
	http.SetCookie(*w, &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   shouldCookieBeSecure,
		HttpOnly: true,
		Path:     "/",
	})
	return nil
}

func SomeUsersVerified() bool {
	users := api.Container.ListUsersUseCase.ListUsers()
	for _, localUser := range users {
		if localUser.IsVerified {
			return true
		}
	}
	return false
}

func IsUserTableEmpty() bool {
	users := api.Container.ListUsersUseCase.ListUsers()
	return len(users) == 0
}

func IsLoggedIn(r *http.Request) bool {
	token, err := jwtauth.VerifyRequest(
		Token,
		r,
		jwtauth.TokenFromCookie,
		jwtauth.TokenFromHeader,
		jwtauth.TokenFromQuery)
	return token != nil && err == nil
}

func IsVerified(r *http.Request) bool {
	token, err := jwtauth.VerifyRequest(
		Token,
		r,
		jwtauth.TokenFromCookie,
		jwtauth.TokenFromHeader,
		jwtauth.TokenFromQuery)
	if err != nil {
		return false
	}
	userId := token.PrivateClaims()["user_id"].(float64)
	currentUser, _ := api.Container.GetUserUseCase.GetUser(uint32(userId))
	return currentUser.IsVerified
}

func GetUserFromCredentials(credentials LoginCredentials) (user.User, error) {
	dbUser, err := api.Container.GetUserUseCase.GetUserByUsername(credentials.Username)
	if err != nil {
		return user.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(credentials.Password))
	if err != nil {
		return user.User{}, err
	}

	return dbUser, nil
}

func GetNewUser(newUserCredentials RegisterCredentials, verificationCode string) (user.User, error) {
	createdUser, err := api.Container.CreateUserUseCase.CreateUser(useCases.CreateUserCommand{
		Username:         newUserCredentials.Username,
		Password:         newUserCredentials.Password,
		Email:            newUserCredentials.Email,
		VerificationCode: verificationCode,
	})
	if err != nil {
		return user.User{}, err
	}
	return createdUser, nil
}

func login(w http.ResponseWriter, r *http.Request) {
	var credentials LoginCredentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, BodyErrorMessage, http.StatusBadRequest)
		return
	}

	err = api.Validate.Struct(credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbUser, err := GetUserFromCredentials(credentials)
	if err != nil {
		http.Error(w, LogsErrorMessage, http.StatusForbidden)
		return
	}

	_ = SetJwtCookie(&w, dbUser.ID)

	message, _ := json.Marshal(map[string]any{"message": "User logged in! HTTPonly jwt cookie created"})
	_, _ = w.Write(message)
}

func register(w http.ResponseWriter, r *http.Request) {
	if !IsLoggedIn(r) && !SomeUsersVerified() {
		http.Error(w, "You are not allowed to create a user. Log in or reset database.", http.StatusForbidden)
		return
	}

	var credentials RegisterCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, BodyErrorMessage, http.StatusBadRequest)
		return
	}

	err = api.Validate.Struct(credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	verificationCode := uuid.NewString()
	createdUser, err := GetNewUser(credentials, verificationCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = SetJwtCookie(&w, createdUser.ID)

	message, _ := json.Marshal(map[string]any{"message": "User registered! HTTPonly jwt cookie created"})
	_, _ = w.Write(message)
}

func RemoveJwtCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now(),
		MaxAge:   -1,
		Secure:   shouldCookieBeSecure,
		HttpOnly: true,
		Path:     "/",
	})
}

func logout(w http.ResponseWriter, _ *http.Request) {
	RemoveJwtCookie(w)
	message, _ := json.Marshal(map[string]any{"message": "User logged out! HTTPonly jwt cookie deleted"})
	_, _ = w.Write(message)
}

func NewAuthRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/register", register)
	r.Post("/login", login)
	r.Post("/logout", logout)
	return r
}
