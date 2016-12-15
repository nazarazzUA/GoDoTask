package cli

import (
    "github.com/fatih/color"
)

type CommandHandler func();

type Command struct {
	name string;
	desc string;
	handler CommandHandler
}

func (com *Command) info () {
	color.Cyan("%-30s : %s \n",com.name, com.desc);
}

func (com *Command) exec () {
	com.handler();
}