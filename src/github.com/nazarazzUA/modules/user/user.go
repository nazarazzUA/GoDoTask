package user

import (
	"github.com/go-martini/martini"
	"github.com/nazarazzUA/modules/user/handlers"
)

type UserModule struct { }

func (mod *UserModule) Config(m *martini.ClassicMartini) {

}

func (mod *UserModule) InitializeHandlers(m *martini.ClassicMartini) {

	m.Get("/login", handlers.LoginPage)
	m.Post("/login",handlers.LoginProcess)

}
