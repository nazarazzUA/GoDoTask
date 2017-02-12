package core

import (
	"github.com/martini-contrib/render"
	_ "fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"github.com/nazarazzUA/modules/core/handlers"
	"github.com/nazarazzUA/modules/core/middleware"
	"github.com/nazarazzUA/app"
)

type CoreModule struct {
	store sessions.CookieStore

}


func (mod *CoreModule) Config(application *app.WebApplication) {

	application.CliApp.RegisterHandler(&DataBaseCommands{});
	application.CliApp.RegisterHandler(&WebCommands{webApp:application});
	application.CliApp.RegisterHandler(&FixtureCommand{});

	mod.store = sessions.NewCookieStore([]byte("LKJAHFfjaadfaw"))
	options := render.Renderer(render.Options{
		Directory: "templates",
		Layout: "layout",
		Extensions: []string{".tmpl", ".html"},
	});

	application.Martini.Use(martini.Static("assets", martini.StaticOptions{ Exclude: "/api/"}));
	application.Martini.Use(options);
	application.Martini.Use(sessions.Sessions("sess_id", mod.store))
	application.Martini.Use(middleware.AuthUser);

	db := GetDb();

	application.Martini.Map(db);

}

func (mod *CoreModule) InitializeHandlers(m *martini.ClassicMartini) {

	m.Get("/", handlers.GetIndexPage)
	m.Get("/admin/**", handlers.ServeAdminStaticFiles)
	m.Get("/404", handlers.NotFoundHandler)
	m.NotFound(handlers.NotFoundHandler);
}
