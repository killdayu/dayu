package main

import (
	"dayu/env"
	"dayu/help"
	"dayu/parameters"
	"github.com/fatih/color"
)

var myenv env.Env
var mypar parameters.Par

func main() {
	color.Yellow("dayu v0.1")
	myenv.Set()
	myenv.Show()

	mypar.Init()
	if mypar.Len == 0 {
		help.MainHelp()
	}
	mypar.Switch(myenv)

}
