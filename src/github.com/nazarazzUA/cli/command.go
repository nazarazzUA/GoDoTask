package cli

import "fmt"

type CommandHandler func();

type Command struct {
	name string;
	desc string;
	handler CommandHandler
}

func (com *Command) info () {
	fmt.Printf("%-30s : %s \n",com.name, com.desc);
}

func (com *Command) exec () {
	com.handler();
}