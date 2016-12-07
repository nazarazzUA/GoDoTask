package core

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
)

type CoreModule struct {

}

func (mod *CoreModule) Config() {


}

func (mod *CoreModule) UseMiddleware (r *httprouter.Router) {


}

func (mod *CoreModule) InitializeHandlers(r *httprouter.Router) {

	r.GET("/", Index)

}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}