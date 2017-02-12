package cli

import (
	"sort"
)

type CliApplication struct {
	arguments []string
	command map[string] *Command;
}

func NewCliApp() *CliApplication {

	appl := &CliApplication{
		command: make(map[string] *Command),
	};

	return appl;
}

func (app *CliApplication) AddHandler(name string, desc string, fn CommandHandler) {
	app.command[name] = &Command{name:name,desc:desc,handler:fn};
}

func (app *CliApplication) ShowAllCommand() {

	var sortedKeys []string;

	for key, _ := range app.command {
		sortedKeys = append(sortedKeys, key);
	}

	sort.Strings(sortedKeys);

	for _, commandName := range sortedKeys {
		app.command[commandName].info();
	}
}

func (app *CliApplication) ExecuteCommand(inCommand []string) {

	if (len(inCommand) == 0) {
		app.ShowAllCommand();
		return;
	}

	if command, ok := app.command[inCommand[0]]; ok {
		command.exec();
		return;
	}

	app.ShowAllCommand();
}

func (app *CliApplication) RegisterHandler (com CommandsRegisterHandler)  {
	com.RegisterCommand(app);
}

func (app *CliApplication) RegisterWebHandler (com CommandsRegisterHandler)  {
	com.RegisterCommand(app);
}