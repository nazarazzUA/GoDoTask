package middleware

import (
	"github.com/martini-contrib/sessions"
	"net/http"
	"github.com/martini-contrib/render"
	//"fmt"
)

const (
	USER_LOGIN_PATH = "/login"
	USER_REGISTER_PATH = "/register"
);

func AuthUser (s sessions.Session , w http.ResponseWriter, r *http.Request, render render.Render )  {

	if r.URL.String() == USER_LOGIN_PATH || r.URL.String() == USER_REGISTER_PATH {
		return;
	}

	if s.Get("user_id") == nil {
		render.Redirect("/login", 302)
		return
	}
}