package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "Devlog"
	app.Usage = "Keep a development / done log from your cli"

	app.Commands = []cli.Command{
		{

			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add a task which has been done",
			Action: func(c *cli.Context) {

				println("Hello World!")
			},
		},
	}

	app.Run(os.Args)

}
