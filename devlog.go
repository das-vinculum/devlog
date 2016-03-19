package main

import (
	"github.com/codegangsta/cli"
	"github.com/das-vinculum/devlog/storage"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
				logEntry := &storage.Logentry{
					Entry: strings.Join(c.Args()[:], " "),
					Date:  time.Now(),
				}
				logEntry.Store()
			},
		},
	}

	app.Run(os.Args)

}
