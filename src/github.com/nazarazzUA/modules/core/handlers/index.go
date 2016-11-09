package handlers

import (
	"io/ioutil"
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"path/filepath"
	"mime"
)

const (
	ADMIN_APP = "GoDoTaskAdminApp"
	ADMIN_APP_PATH = "./admin_app/templates/index.html"
);

func readAdminIndexFIle (w http.ResponseWriter, r * http.Request) {
	body, err := ioutil.ReadFile(ADMIN_APP_PATH)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type","text/html")
	w.Write(body);
}

func GetIndexPage(w http.ResponseWriter, r * http.Request) {
	readAdminIndexFIle(w,r);
}

func ServeAdminStaticFiles(w http.ResponseWriter, r render.Render, params martini.Params, req *http.Request) {

	fileName  := params["_1"]
	body, err := ioutil.ReadFile("./admin_app/" + fileName)

	if err != nil {
		if req.Header.Get("X-Requested-With") == ADMIN_APP {
			w.WriteHeader(404);
			w.Write([]byte("{\"error\":\"Not found template\"}"))
			return;
		}
		r.Redirect("/404")
	}

	ext   := filepath.Ext(fileName)
	mType := mime.TypeByExtension(ext)

	w.Header().Set("Content-Type",mType);
	w.Write(body);
}

func NotFoundHandler (w http.ResponseWriter, r * http.Request) {
	readAdminIndexFIle(w,r);
}