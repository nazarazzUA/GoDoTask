package core

import (

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/nazarazzUA/modules/core/handlers"
	"github.com/nazarazzUA/modules/core/middlewares"
)

type CoreModule struct {
	store sessions.CookieStore
}

func (mod *CoreModule) Config() {
	mod.store = sessions.NewCookieStore([]byte("LKJAHFfjaadfaw"))
}

func (mod *CoreModule) UseMiddleware (m *martini.ClassicMartini) {

	options := render.Options{
		Directory: "templates",
		Layout: "layout",
		Extensions: []string{".tmpl"},
		IndentJSON: true,
	}

	m.Use(sessions.Sessions("GoDoTask", mod.store))
	m.Use(martini.Static("assets", martini.StaticOptions{Exclude: "/admin"}))
	m.Use(render.Renderer(options))
	m.Use(middlewares.CheckUserSessions);

}

func (mod *CoreModule) InitializeHandlers(m *martini.ClassicMartini) {

	m.Get("/", handlers.GetIndexPage)
	m.Get("/admin/**", handlers.ServeAdminStaticFiles)
	m.Get("/404", handlers.NotFoundHandler)
	m.NotFound(handlers.NotFoundHandler);
}
