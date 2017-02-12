package core

import (
	"github.com/nazarazzUA/app/cli"
	"fmt"
	"github.com/nazarazzUA/models"

)

type DataBaseCommands struct {}

func (dbComm * DataBaseCommands) RegisterCommand(cli *cli.CliApplication) {
	cli.AddHandler("db.drop.table", "drop application database", dbComm.dropDb);
	cli.AddHandler("db.migrate", "Create migration in database", dbComm.migration);
}

func (dbComm * DataBaseCommands)  dropDb() {

	fmt.Println("Drop all tabels data");

	db := GetDb();
	db.DropTableIfExists(&models.User{});
	defer db.Close()

	fmt.Println("Table dropped successfull!");
}

func (dbComm * DataBaseCommands)  migration() {

	fmt.Println("Start Create Migrations");

	db := GetDb();
	db.AutoMigrate(&models.User{});

	defer db.Close()

	fmt.Println("Migration created successfull!");
}