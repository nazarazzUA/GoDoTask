package app

import "github.com/go-martini/martini"

type Module interface {

	Config()
	UseMiddleware(m *martini.ClassicMartini)
	InitializeHandlers(m *martini.ClassicMartini)
}
