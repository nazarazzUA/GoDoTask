package handlers

import "github.com/martini-contrib/render"

func LoginPage(r render.Render) {
	r.HTML(200,"default/users/login",nil);
}
