package app

import (
	"github.com/go-martini/martini"
)

type Module interface {
	Config(application *WebApplication)
	InitializeHandlers(m *martini.ClassicMartini)
}
