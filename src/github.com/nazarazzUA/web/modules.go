package web

import (
	"github.com/julienschmidt/httprouter"
)

type Module interface {

	Config()
	UseMiddleware(r *httprouter.Router)
	InitializeHandlers(r *httprouter.Router)
}
