package app

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/nazarazzUA/app/cli"
)

type WebApplication struct {

	Martini *martini.ClassicMartini
	CliApp *cli.CliApplication
	modules []Module
}

func (app *WebApplication) RunModules() {
	for _,module := range app.modules {
		module.Config(app)
	    module.InitializeHandlers(app.Martini)
	}
}

func (app *WebApplication) SetModule(m Module) {
	app.modules = append(app.modules, m);
}

func CreateApp () (*WebApplication) {
	return &WebApplication{
		Martini : martini.Classic(),
		CliApp : cli.NewCliApp(),
	}
}

func (app *WebApplication) Run () {
	fmt.Println("Server start listen on port 8080");
	log.Fatal(http.ListenAndServe(":8080", app.Martini));
}

func (app *WebApplication) ExecuteCommand(inCommand []string) {
	app.CliApp.ExecuteCommand(inCommand);
}

