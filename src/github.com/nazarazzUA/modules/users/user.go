package users

import (
	"github.com/go-martini/martini"
	"github.com/nazarazzUA/modules/core"
	"github.com/nazarazzUA/modules/users/handlers"
)

type UserModule struct {
	core.CoreModule
}


func (mod *UserModule) InitializeHandlers(m *martini.ClassicMartini) {

	m.Get("/login",handlers.GetLoginPage)
	m.Post("/login", handlers.LoginProcess)
	m.Get("/registration", handlers.RegisterPage)
	m.Post("/registration", handlers.RegisterProcess)
	m.Any("/logout", handlers.LogoutUser)

}

