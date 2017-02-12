package core

import (
	"fmt"
	"github.com/nazarazzUA/app"
	"github.com/nazarazzUA/app/cli"
)

type WebCommands struct {
	webApp *app.WebApplication
}

func (web * WebCommands) RegisterCommand(cli *cli.CliApplication) {

	handler := func (){
		web.webApp.Run();
	}

	cli.AddHandler("web.start-server", "start golang dev server as port 8080", handler);
	cli.AddHandler("web.stop-server", "stop golang dev server as port 8080", web.stopWebServer);
}

func (web * WebCommands) stopWebServer() {
	fmt.Println("Stop web server");
}


