package core

import (
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/nazarazzUA/modules/core/handlers"
)

type CoreModule struct {
	store sessions.CookieStore
}

func (mod *CoreModule) Config() {

	mod.store = sessions.NewCookieStore([]byte("LKJAHFfjaadfaw"))
}

func (mod *CoreModule) UseMiddleware (m *martini.ClassicMartini) {

	m.Use(sessions.Sessions("GoDoTask", mod.store))
	m.Use(martini.Static("assets", martini.StaticOptions{Exclude: "/admin"}))

	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout: "layout",
		Extensions: []string{".tmpl"},
		IndentJSON: true,
	}))

	m.Use(func (s sessions.Session, r render.Render, req *http.Request) {

		if req.URL.String() == "/login" || req.URL.String() == "/registration" {
			return;
		}

		if s.Get("user_id") == nil {
			r.Redirect("/login", 302)
			return
		}
	});
}

func (mod *CoreModule) InitializeHandlers(m *martini.ClassicMartini) {

	m.Get("/", handlers.GetIndexPage)
	m.Get("/admin/**", handlers.ServeAdminStaticFiles)
	m.Get("/404", handlers.NotFoundHandler)
	m.NotFound(handlers.NotFoundHandler);
}