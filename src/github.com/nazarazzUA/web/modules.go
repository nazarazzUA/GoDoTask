package web

import (
	"github.com/go-martini/martini"
)

type Module interface {

	Config(m *martini.ClassicMartini)
	InitializeHandlers(m *martini.ClassicMartini)
}
