package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/nazarazzUA/app"
	"github.com/nazarazzUA/modules/users/models"
)

func GetLoginPage(r render.Render) {
	r.HTML(200, "default/users/login", nil)
}

func LoginProcess (r render.Render, s sessions.Session , req *http.Request) {

	fmt.Println(s.Get("user_id"));

	email := req.FormValue("email");
	pass  := req.FormValue("password");

	fmt.Println(email);
	fmt.Println(pass);

	r.HTML(200, "default/users/login", nil)
}


func RegisterPage (r render.Render, s sessions.Session , req *http.Request) {

	r.HTML(200, "default/users/registration", nil)
}


func RegisterProcess (r render.Render, s sessions.Session , req *http.Request) {

	email := req.FormValue("email");
	pass  := req.FormValue("password");

	if messages, ok  := isDataValid(email, pass); !ok {
		r.HTML(403, "default/users/registration", messages)
		return
	}

	password := []byte(pass)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		r.HTML(200, "default/errors/500", nil)
		return
	}

	db, err := app.GetDb()

	if err != nil {
		r.HTML(200, "default/errors/500", nil)
		return
	}

	db.Create(&models.User{Email: email,Password: string(hashedPassword)});

	r.Redirect("/login");
}


func isDataValid(email string, password string) (*models.ValidationErrorMessages, bool) {

	messages := &models.ValidationErrorMessages{};

	if len(strings.TrimSpace(email)) == 0 {
		messages.Email = "This field can`t bee empty!";
		return messages, false;
	}

	if len(strings.TrimSpace(password)) < 3 {
		messages.Password = "Password to short!";
		return messages, false;
	}

	if len(strings.TrimSpace(password)) > 18 {
		messages.Password = "Password to long. Your don`t remember this shit!"
		return messages, false;
	}


	return messages, true;

}