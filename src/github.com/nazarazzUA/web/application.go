package web

import (
	"github.com/julienschmidt/httprouter"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type WebApplication struct {

	r *httprouter.Router
	modules []Module
}

func (app *WebApplication) runModules() {
	for _,module := range app.modules {
		module.Config()
		module.UseMiddleware(app.r)
		module.InitializeHandlers(app.r)
	}
}

func (app *WebApplication) SetModule(m Module) {
	app.modules = append(app.modules, m);
}

func CreateApp () (*WebApplication) {
	return &WebApplication{
		r : httprouter.New(),
	}
}

func (app *WebApplication) Run () {

	app.runModules();

}
