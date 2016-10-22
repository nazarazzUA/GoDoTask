package main

import (
	"github.com/nazarazzUA/app"
	"github.com/nazarazzUA/modules"
)

func main() {

	application := app.CreateApp();
	application.SetModule(&modules.CoreModule{});
	application.Run();

}




