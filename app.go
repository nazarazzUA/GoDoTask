package main

import (
	"os"
	"github.com/nazarazzUA/cli"
	"github.com/nazarazzUA/commands"
)

func main() {

	cliApp := cli.NewCliApp();
	commands.RegisterAllCommand(cliApp);
	cliApp.ExecuteCommand(os.Args[1:]);

}




