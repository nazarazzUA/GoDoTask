package main

import (

	"github.com/nazarazzUA/cli"
	"github.com/nazarazzUA/commands"
	"os"
)

func main() {

	cliApp := cli.NewCliApp();
	commands.RegisterAllCommand(cliApp);
	cliApp.ExecuteCommand(os.Args[1:]);

}




