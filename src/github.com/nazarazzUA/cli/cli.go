package cli


type CliApplication struct {
	arguments []string
	command map[string] *Command
}

func NewCliApp() *CliApplication {

	app := &CliApplication{
		command:make(map[string]*Command),
	};

	return app;
}

func (app *CliApplication) AddHandler(name string, desc string, fn CommandHandler) {
	app.command[name] = &Command{name:name,desc:desc,handler:fn};
}

func (app *CliApplication) ShowAllCommand() {

	for _, command := range app.command {
		command.info();
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