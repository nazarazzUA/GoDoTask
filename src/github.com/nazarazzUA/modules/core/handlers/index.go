package handlers

import (
	"io/ioutil"
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"path/filepath"
	"mime"
	"fmt"
)

const ADMIN_APP = "GoDoTaskAdminApp";

func GetIndexPage(w http.ResponseWriter, r * http.Request) {
	body, err := ioutil.ReadFile("./admin_app/templates/index.html")
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type","text/html")
	w.Write(body);
}

func ServeAdminStaticFiles(w http.ResponseWriter, r render.Render, params martini.Params, req *http.Request) {

	fileName := params["_1"]
	body, err := ioutil.ReadFile("./admin_app/" + fileName)

	if err != nil {
		if req.Header.Get("X-Requested-With") == ADMIN_APP {
			w.WriteHeader(404);
			return;
		}
		r.Redirect("/404")
	}

	ext := filepath.Ext(fileName)
	mType := mime.TypeByExtension(ext)

	w.Header().Set("Content-Type",mType);
	w.Write(body);
}

func NotFoundHandler (r render.Render, req * http.Request) {
	fmt.Println();

	r.HTML(404, "default/errors/404", nil);
}