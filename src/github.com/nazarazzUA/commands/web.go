package commands

import (
	"fmt"
	"github.com/nazarazzUA/cli"
	"github.com/nazarazzUA/web"
)

type WebCommands struct {}

func (web * WebCommands) RegisterCommand(cli *cli.CliApplication) {

	cli.AddHandler("web.start-server", "start golang dev server as port 8080", web.startServer);
	cli.AddHandler("web.stop-server", "stop golang dev server as port 8080", web.stopWebServer);
}

func (w * WebCommands) startServer() {
	web.StartWebApp();
}

func (web * WebCommands) stopWebServer() {
	fmt.Println("Stop web server");
}


