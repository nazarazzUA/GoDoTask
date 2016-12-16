package handlers

import (
	"github.com/martini-contrib/render"
	"net/http"
	"github.com/nazarazzUA/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/nazarazzUA/modules/core"
	"github.com/martini-contrib/sessions"
)

const EMPTY_STRING = "";

func LoginPage(r render.Render) {
	r.HTML(200,"default/users/login",nil);
}

func LoginProcess(s sessions.Session, req * http.Request, r render.Render) {

	email := req.FormValue("email");
	pass  := req.FormValue("password");

	user, errorMessage := loginProcess(email, pass);
	if errorMessage != EMPTY_STRING {
		data := struct { Message string }{errorMessage};
		r.HTML(400,"default/users/login", data);
		return;
	}

	s.Set("user_id",user.ID);
	r.Redirect("/")
}

func loginProcess(email, password string) (*models.User, string) {

	var errorMessage string;
	user := &models.User{};
	if email == EMPTY_STRING || password == EMPTY_STRING {
		errorMessage = "Email or pasword cant be empty";
		return user, errorMessage;
	};

	db := core.GetDb();
	db.Where(&models.User{Email:email}).First(&user)

	if user != nil {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return nil, "Wrong email or password"
		}
	}
	return user, EMPTY_STRING;
}