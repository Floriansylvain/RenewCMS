package pages

import (
	"RenewCMS/api"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var PasswordLinkErrorMessage = "The reset password link you used is invalid."
var PasswordLinkSuccessMessage = "Your account password was successfully updated."

func GetPasswordResetValidate(w http.ResponseWriter, r *http.Request) {
	failure := r.URL.Query().Get("failure")
	pageError := NewPageError("")
	if failure != "" {
		pageError = NewPageError(failure)
	}
	template, _ := api.Container.GetPageUseCase.GetPage("passwordResetValidate", map[string]any{
		"Head":  headTmpl,
		"Error": pageError,
		"Email": r.URL.Query().Get("email"),
		"Code":  r.URL.Query().Get("c"),
	})
	_, _ = w.Write(template)
}

func PostPasswordResetValidate(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	resetCode := r.FormValue("code")

	var redirectionErrorLink = "/register/reset/validate?email=" + email + "&c=" + resetCode

	if len(password) < 8 {
		http.Redirect(w, r, redirectionErrorLink+"&failure=Password length should be at least 8 characters.", http.StatusSeeOther)
		return
	}

	fetchedUser, err := api.Container.GetUserUseCase.GetUserByEmail(email)
	if err != nil {
		http.Redirect(w, r, redirectionErrorLink+"&failure="+PasswordLinkErrorMessage, http.StatusSeeOther)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(fetchedUser.PasswordResetCode), []byte(resetCode))
	if err != nil {
		http.Redirect(w, r, redirectionErrorLink+"&failure="+PasswordLinkErrorMessage, http.StatusSeeOther)
		return
	}

	_, _ = api.Container.UpdateUserUseCase.UpdatePassword(fetchedUser.ID, password)
	_, _ = api.Container.UpdateUserUseCase.UpdatePasswordResetCode(fetchedUser.ID, "")

	http.Redirect(w, r, "/login?success="+PasswordLinkSuccessMessage, http.StatusSeeOther)
}
