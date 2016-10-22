package modules

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/nazarazzUA/modules/handlers"
)

type CoreModule struct {
	store sessions.CookieStore
}

func (mod *CoreModule) Config() {

	mod.store = sessions.NewCookieStore([]byte("LKJAHFfjaadfaw"))
}

func (mod *CoreModule) UseMiddleware (m *martini.ClassicMartini) {

	m.Use(sessions.Sessions("GoDoTask", mod.store))
	m.Use(martini.Static("assets", martini.StaticOptions{Exclude: "/api"}))

	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout: "layout",
		Extensions: []string{".tmpl"},
		IndentJSON: true,
	}))
}

func (mod *CoreModule) InitializeRoute(m *martini.ClassicMartini) {

}

func (mod *CoreModule) InitializeHandlers(m *martini.ClassicMartini) {
	m.Get("/", handlers.GetIndexPage)
}