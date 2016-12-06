package commands

import (
	"github.com/nazarazzUA/cli"
	"fmt"
)

type DataBaseCommands struct {}

func (dbComm * DataBaseCommands) RegisterCommand(cli *cli.CliApplication) {
	cli.AddHandler("db.create", "create application database", dbComm.createDb);
	cli.AddHandler("db.drop-data-base","drop application database", dbComm.dropDb);
	cli.AddHandler("db.migration-create","Create migration in database", dbComm.migration);
}

func (dbComm * DataBaseCommands)  createDb() {
	fmt.Println("Hello from create db");
}

func (dbComm * DataBaseCommands)  dropDb() {
	fmt.Println("Drop data base");
}

func (dbComm * DataBaseCommands)  migration() {
	fmt.Println("Create Migrations");
}