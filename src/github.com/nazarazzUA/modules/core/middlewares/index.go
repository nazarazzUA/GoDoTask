package middlewares

import (
	"github.com/martini-contrib/render"
	"net/http"
	"github.com/martini-contrib/sessions"
)

func CheckUserSessions(s sessions.Session, r render.Render, req *http.Request) {

	if req.URL.String() == "/login" || req.URL.String() == "/registration" {
		return;
	}

	if s.Get("user_id") == nil {
		r.Redirect("/login", 302)
		return
	}
}
