package commands

import "github.com/nazarazzUA/cli"

type CommandsRegisterHandler interface {
	RegisterCommand(cli *cli.CliApplication);
}

var commands []CommandsRegisterHandler;

func RegisterAllCommand(cli *cli.CliApplication) {

	commands = []CommandsRegisterHandler{
		&DataBaseCommands{},
		&WebCommands{},
		&FixtureCommand{},
	}

	for _, com := range commands {
		com.RegisterCommand(cli);
	}
}


