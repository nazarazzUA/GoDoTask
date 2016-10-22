package app

import (
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Application struct {

	m *martini.ClassicMartini
	modules []Module
}

func (app *Application) runModules() {

	for _,module := range app.modules {
		module.Config()
		module.UseMiddleware(app.m)
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

func GetDb () (*gorm.DB, error) {

	db, err := gorm.Open("mysql", "root:root@/godutask?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return db, nil;
}