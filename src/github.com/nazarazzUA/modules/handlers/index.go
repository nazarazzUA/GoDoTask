package handlers

import "github.com/martini-contrib/render"

func GetIndexPage(r render.Render) {
	r.HTML(200, "default/main-page", nil)
}
