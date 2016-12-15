package main

import (
	//"os"
	//"github.com/nazarazzUA/cli"
	//"github.com/nazarazzUA/commands"
	//"golang.org/x/crypto/bcrypt"
	//"fmt"
	"github.com/nazarazzUA/cli"
	"github.com/nazarazzUA/commands"
	"os"
)

func main() {

	//password := []byte("111111")
	//
	//// Hashing the password with the default cost of 10
	//hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(hashedPassword))

	//// Comparing the password with the hash
	//err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	//fmt.Println(err) // nil means it is a match
	//
	cliApp := cli.NewCliApp();
	commands.RegisterAllCommand(cliApp);
	cliApp.ExecuteCommand(os.Args[1:]);

}




