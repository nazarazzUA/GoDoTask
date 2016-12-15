package web

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nazarazzUA/modules/core"
	"log"
	"net/http"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/nazarazzUA/modules/user"
)

type WebApplication struct {

	m *martini.ClassicMartini
	modules []Module
}

func (app *WebApplication) runModules() {
	for _,module := range app.modules {
		module.Config(app.m)
	    module.InitializeHandlers(app.m)
	}
}

func (app *WebApplication) SetModule(m Module) {
	app.modules = append(app.modules, m);
}

func CreateApp () (*WebApplication) {
	return &WebApplication{
		m : martini.Classic(),
	}
}

func (app *WebApplication) Run () {
	app.runModules();
	fmt.Println("Server start listen on port 8080");
	log.Fatal(http.ListenAndServe(":8080", app.m));
}

func StartWebApp() {

	app := CreateApp();
	app.SetModule(&core.CoreModule{})
	app.SetModule(&user.UserModule{})
	app.Run();

}