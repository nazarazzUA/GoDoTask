package main

import (
	"os"
	"github.com/nazarazzUA/modules/core"
	"github.com/nazarazzUA/modules/user"
	"github.com/nazarazzUA/app"
)

func main() {

	application := app.CreateApp();
	application.SetModule(&core.CoreModule{})
	application.SetModule(&user.UserModule{})
	application.RunModules();

	application.ExecuteCommand(os.Args[1:]);

}




