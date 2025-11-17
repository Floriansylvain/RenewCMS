package pages

import (
	"RenewCMS/api"
	"RenewCMS/api/controllers/auth"
	"net/http"
	"net/url"
)

const LoginRoute = "/login"

type LoginPage struct {
	PageError *PageError `json:"error"`
	Username  string     `json:"username"`
}

var EmptyLoginPage = &LoginPage{
	PageError: NewPageError(""),
	Username:  "",
}

func GetLoginPageHandler(loginPage *LoginPage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if auth.IsLoggedIn(r) {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
		if r.Method == http.MethodPost {
			PostLoginPage(w, r)
			return
		}
		if auth.IsUserTableEmpty() {
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}
		success, _ := url.QueryUnescape(r.URL.Query().Get("success"))
		failure, _ := url.QueryUnescape(r.URL.Query().Get("failure"))
		bs, _ := api.Container.GetPageUseCase.GetPage("login", map[string]any{
			"PageError": loginPage.PageError,
			"Username":  loginPage.Username,
			"Head":      headTmpl,
			"Success":   success,
			"Failure":   failure,
		})
		_, _ = w.Write(bs)
	}
}

func PostLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	credentials := auth.LoginCredentials{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	err := api.Validate.Struct(credentials)
	if err != nil {
		r.Method = http.MethodGet
		GetLoginPageHandler(&LoginPage{
			PageError: NewPageError("Invalid username or password format."),
			Username:  r.FormValue("username"),
		})(w, r)
		return
	}

	dbUser, err := auth.GetUserFromCredentials(credentials)
	if err != nil {
		r.Method = http.MethodGet
		GetLoginPageHandler(&LoginPage{
			PageError: NewPageError("Invalid username or password combination."),
			Username:  r.FormValue("username"),
		})(w, r)
		return
	}

	_ = auth.SetJwtCookie(&w, dbUser.ID)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
