package handlers

import (
	"net/http"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/nazarazzUA/app"
	"github.com/nazarazzUA/modules/users/problem"
	"github.com/nazarazzUA/modules/users/models"
)

func GetLoginPage(r render.Render) {
	r.HTML(200, "default/users/login", nil)
}

func LoginProcess (r render.Render, s sessions.Session , req *http.Request) {

	email := req.FormValue("email");
	pass  := req.FormValue("password");

	if messages, ok  := isDataValid(email, pass); !ok {
		r.HTML(403, "default/users/login", messages)
		return
	}

	db, err := app.GetDb()
	defer db.Close();

	if err != nil {
		r.HTML(200, "default/errors/500", nil)
		return
	}

	findsUser := &models.User{}
	db.Where(&models.User{Email:email}).First(&findsUser)
	err = bcrypt.CompareHashAndPassword([]byte(findsUser.Password), []byte(pass))

	if err != nil {
		r.HTML(403, "default/users/login", &problem.ValidationErrorMessages{Email:"User Email or Password Incorect!"})
		return
	}

	s.Set("user_id", findsUser.ID);
	r.Redirect("/");
}


func RegisterPage (r render.Render, s sessions.Session , req *http.Request) {

	r.HTML(200, "default/users/registration", nil)
}

func LogoutUser(r render.Render, s sessions.Session) {
	s.Set("user_id", nil);
	r.Redirect("/login");
}

func RegisterProcess (r render.Render, req *http.Request) {

	email := req.FormValue("email");
	pass  := req.FormValue("password");

	if messages, ok  := isDataValid(email, pass); !ok {
		r.HTML(403, "default/users/registration", messages)
		return
	}

	hashedPassword := cryptPassword(pass)

	db, err := app.GetDb()
	defer db.Close();

	if err != nil {
		r.HTML(200, "default/errors/500", nil)
		return
	}

	user := &models.User{};
	db.Where(&models.User{Email:email}).First(&user);

	if user.Email == email {
		r.HTML(403, "default/users/registration",
			&problem.ValidationErrorMessages{
				Email:"With this email user exists!",
			})
		return
	}

	db.Create(&models.User{Email: email, Password: string(hashedPassword)});
	r.Redirect("/login");
}


func isDataValid(email string, password string) (*problem.ValidationErrorMessages, bool) {

	messages := &problem.ValidationErrorMessages{};

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

func cryptPassword (pass string) ([]byte) {

	password := []byte(pass)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err);
	}

	return hashedPassword
}