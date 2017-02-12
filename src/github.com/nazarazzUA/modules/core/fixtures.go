package core

import (
	"github.com/nazarazzUA/app/cli"
	"github.com/nazarazzUA/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/fatih/color"
)

type FixtureCommand struct {}

func (fix * FixtureCommand) RegisterCommand(cli *cli.CliApplication) {

	cli.AddHandler("fix.load-user", "Create test user for app login:nazarazz@bigmir.net password: 111111 ",
		fix.createTestUser);
}

func (w * FixtureCommand) createTestUser() {

	password := []byte("111111")
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	user := models.User{
		Email:"nazarazz@bigmir.net",
		Password: string(hashedPassword),
		FirstName: "Nazar",
		LastName: "Shchepilow",
	}

	db := GetDb();
	db.Create(&user);

	color.Green("User is created sucsesfull");
}
