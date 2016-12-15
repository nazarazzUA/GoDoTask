package core

import (
	"github.com/martini-contrib/render"
	_ "fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"github.com/nazarazzUA/modules/core/handlers"
	"github.com/nazarazzUA/modules/core/middleware"
)

type CoreModule struct {
	store sessions.CookieStore
}


func (mod *CoreModule) Config(m *martini.ClassicMartini) {

	mod.store = sessions.NewCookieStore([]byte("LKJAHFfjaadfaw"))
	options := render.Renderer(render.Options{
		Directory: "templates",
		Layout: "layout",
		Extensions: []string{".tmpl", ".html"},
	});

	m.Use(martini.Static("assets", martini.StaticOptions{ Exclude: "/api/"}));
	m.Use(options);
	m.Use(sessions.Sessions("sess_id", mod.store))
	m.Use(middleware.AuthUser);

	db := GetDb();

	m.Map(db);

}

func (mod *CoreModule) InitializeHandlers(m *martini.ClassicMartini) {

	m.Get("/", handlers.GetIndexPage)
	m.Get("/admin/**", handlers.ServeAdminStaticFiles)
	m.Get("/404", handlers.NotFoundHandler)
	m.NotFound(handlers.NotFoundHandler);
}
