package app

import (
	"github.com/go-martini/martini"
)

type Application struct {

	m *martini.ClassicMartini
	modules []Module
}

func (app *Application) runModules() {

	for _,module := range app.modules {
		module.Config()
		module.UseMiddleware(app.m)
		module.InitializeRoute(app.m)
		module.InitializeHandlers(app.m)
	}
}

func (app *Application) SetModule(m Module) {
	app.modules = append(app.modules, m);
}

func CreateApp () (*Application) {
	return &Application{
		m : martini.Classic(),
	}
}

func (app *Application) Run () {

	app.runModules();
	app.m.Run();
}

