package commands

import (
	"fmt"
	"github.com/nazarazzUA/cli"
)

type WebCommands struct {}

func (web * WebCommands) RegisterCommand(cli *cli.CliApplication) {

	cli.AddHandler("web.start-server", "start golang dev server as port 8000", web.startServer);
	cli.AddHandler("web.stor-server", "stop golang dev server as port 8000", web.stopWebServer);
}

func (web * WebCommands) startServer() {
	fmt.Println("Start web server");
}

func (web * WebCommands) stopWebServer() {
	fmt.Println("Stop web server");
}


