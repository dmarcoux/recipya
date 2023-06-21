package server

import (
	"github.com/google/uuid"
	"github.com/reaper47/recipya/internal/app"
	"github.com/reaper47/recipya/internal/auth"
	"github.com/reaper47/recipya/internal/templates"
	"github.com/reaper47/recipya/internal/utils/regex"
	"golang.org/x/exp/maps"
	"net/http"
	"strings"
	"time"
)

func (s *Server) confirmHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	userID, err := auth.ParseToken(token)
	if err != nil {
		sendErrorAdminEmail(s.Email.Send, "confirmHandler.ParseToken: "+token, err)
		w.WriteHeader(http.StatusBadRequest)
		templates.Render(w, templates.Simple, templates.ErrorTokenExpired)
		return
	}

	if err := s.Repository.Confirm(userID); err != nil {
		sendErrorAdminEmail(s.Email.Send, "confirmHandler.Confirm: "+token, err)
		w.WriteHeader(http.StatusNotFound)
		templates.Render(w, templates.Simple, templates.ErrorConfirm)
		return
	}

	templates.Render(w, templates.Simple, templates.SuccessConfirm)
}

func loginHandler(w http.ResponseWriter, _ *http.Request) {
	page := templates.LoginPage
	templates.Render(w, page, templates.Data{Title: page.Title()})
}

func (s *Server) loginPostHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	if !regex.Email.MatchString(email) || password == "" {
		w.Header().Set("HX-Trigger", makeToast("Credentials are invalid.", errorToast))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	userID := s.Repository.VerifyLogin(email, password)
	if userID == -1 {
		w.Header().Set("HX-Trigger", makeToast("Credentials are invalid.", errorToast))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	sid := uuid.New()
	SessionData[sid] = userID
	http.SetCookie(w, NewSessionCookie(sid.String()))

	if r.FormValue("remember-me") == "true" {
		selector, validator := auth.GenerateSelectorAndValidator()
		http.SetCookie(w, NewRememberMeCookie(selector, validator))
		if err := s.Repository.AddAuthToken(selector, validator, userID); err != nil {
			sendErrorAdminEmail(s.Email.Send, "loginPostHandler.AddAuthToken", err)
		}
	}

	redirectUri := "/"
	if c, err := r.Cookie(cookieNameRedirect); err != http.ErrNoCookie {
		redirectUri = c.Value
	}

	w.Header().Set("HX-Redirect", redirectUri)
	w.WriteHeader(http.StatusSeeOther)
}

func registerHandler(w http.ResponseWriter, _ *http.Request) {
	page := templates.RegisterPage
	templates.Render(w, page, templates.Data{Title: page.Title()})
}

func (s *Server) registerPostEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	if !regex.Email.MatchString(email) {
		templates.RenderComponent(w, "registration", "email-invalid", templates.RegisterData{Email: email})
		return
	}

	if s.Repository.IsUserExist(email) {
		templates.RenderComponent(w, "registration", "email-taken", templates.RegisterData{Email: email})
		return
	}

	templates.RenderComponent(w, "registration", "email-valid", templates.RegisterData{Email: email})
}

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromSessionCookie(r)
	if userID == -1 {
		userID = getUserIDFromRememberMeCookie(r, s.Repository.GetAuthToken)
	}

	if userID == -1 {
		return
	}

	sessionCookie, err := r.Cookie(cookieNameSession)
	if err != http.ErrNoCookie {
		sessionCookie.MaxAge = -1
		http.SetCookie(w, sessionCookie)
	}
	maps.DeleteFunc(SessionData, func(_ uuid.UUID, id int64) bool { return id == userID })

	rememberMeCookie, err := r.Cookie(cookieNameRememberMe)
	if err != http.ErrNoCookie {
		if err := s.Repository.DeleteAuthToken(userID); err != nil {
			sendErrorAdminEmail(s.Email.Send, "logoutHandler.DeleteAuthToken", err)
		}

		rememberMeCookie.MaxAge = -1
		http.SetCookie(w, rememberMeCookie)
	}

	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func (s *Server) registerPostPasswordHandler(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password")
	password2 := r.FormValue("password-confirm")
	if password != password2 {
		templates.RenderComponent(w, "registration", "password-invalid", templates.RegisterData{Email: r.FormValue("email")})
		return
	}

	templates.RenderComponent(w, "registration", "password-valid", templates.RegisterData{
		Email:           r.FormValue("email"),
		PasswordConfirm: password,
	})
}

func (s *Server) registerPostHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if !regex.Email.MatchString(email) || password != r.FormValue("password-confirm") {
		w.Header().Set("HX-Trigger", makeToast("Registration failed.", errorToast))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	hashPassword, err := auth.HashPassword(password)
	if err != nil {
		w.Header().Set("HX-Trigger", makeToast("Error encoding your password.", errorToast))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	userID, err := s.Repository.Register(email, hashPassword)
	if err != nil {
		w.Header().Set("HX-Trigger", makeToast("Registration failed.", errorToast))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	token, err := auth.CreateToken(map[string]any{"userID": userID}, 14*24*time.Hour)
	if err != nil {
		sendErrorAdminEmail(s.Email.Send, "registerPostHandler.CreateToken", err)
		w.Header().Set("HX-Trigger", makeToast("Registration failed.", errorToast))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	s.Email.Send(email, templates.EmailIntro, templates.EmailData{
		Token:    token,
		UserName: strings.Split(email, "@")[0],
		URL:      app.Config.URL,
	})
	w.Header().Set("HX-Redirect", "/auth/login")
	w.WriteHeader(http.StatusSeeOther)
}