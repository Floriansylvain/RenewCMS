package pages

import (
	"RenewCMS/api"
	"RenewCMS/api/controllers/auth"
	"log"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

func PostRegisterPendingPage(w http.ResponseWriter, r *http.Request) {
	token, _ := jwtauth.VerifyRequest(
		auth.Token, r,
		jwtauth.TokenFromCookie,
		jwtauth.TokenFromHeader,
		jwtauth.TokenFromQuery)
	userId := token.PrivateClaims()["user_id"].(float64)
	err := api.Container.DeleteUserUseCase.DeleteUser(uint32(userId))
	if err != nil {
		log.Println(err)
		r.Method = http.MethodGet
		http.Redirect(w, r, "/register/pending", http.StatusSeeOther)
	}

	auth.RemoveJwtCookie(w)

	http.Redirect(w, r, "/register", http.StatusSeeOther)
}

func GetRegisterPendingPage(w http.ResponseWriter, _ *http.Request) {
	registerPendingTmpl, _ := api.Container.GetPageUseCase.GetPage("registerPending", map[string]any{
		"Head": headTmpl,
	})
	_, _ = w.Write(registerPendingTmpl)
}
