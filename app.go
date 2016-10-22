package main

import (
	"github.com/nazarazzUA/app"
	"github.com/nazarazzUA/modules/core"
	"github.com/nazarazzUA/modules/users"
)

func main() {

	application := app.CreateApp();
	application.SetModule(&core.CoreModule{});
	application.SetModule(&users.UserModule{});
	application.Run();

}




