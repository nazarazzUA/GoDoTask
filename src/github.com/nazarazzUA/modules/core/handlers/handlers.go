package handlers;

import (
	"github.com/martini-contrib/render"
	"path/filepath"
	"mime"
	"net/http"
	"io/ioutil"
	"github.com/go-martini/martini"
	"fmt"
)

const (
	ADMIN_APP = "GoDoTaskAdminApp"
	ADMIN_APP_PATH = "./admin-app/templates/index.html"
);

var adminAppFile []byte;

func init () {
	var err error;
	adminAppFile, err = ioutil.ReadFile(ADMIN_APP_PATH);
	if err != nil {
		fmt.Println("Admin file not exist");
	}
}

func readAdminIndexFIle (w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type","text/html")
	w.Write(adminAppFile);
}

func ServeAdminStaticFiles(w http.ResponseWriter, r render.Render, params martini.Params, req *http.Request) {

	fileName := params["_1"]
	body, err := ioutil.ReadFile("./admin-app/" + fileName)

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

func GetIndexPage(w http.ResponseWriter, r * http.Request) {
	readAdminIndexFIle(w,r);
}
