package user

import (
	"github.com/go-martini/martini"
	"github.com/nazarazzUA/modules/user/handlers"
	"github.com/nazarazzUA/app"
)

type UserModule struct { }

func (mod *UserModule) Config(application *app.WebApplication) {

}

func (mod *UserModule) InitializeHandlers(m *martini.ClassicMartini) {

	m.Get("/login", handlers.LoginPage)
	m.Post("/login",handlers.LoginProcess)

}
