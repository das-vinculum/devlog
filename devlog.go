package main

import (
	"fmt"
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
				now := time.Now()
				logEntry := &storage.Logentry{
					Entry: strings.Join(c.Args()[:], " "),
					Date:  &now,
				}
				logEntry.Store()
			},
		},
		{
			Name:    "list",
			Aliases: []string{"a"},
			Usage:   "list all done tasks",
			Action: func(c *cli.Context) {
				entries := storage.LoadAllEntries()
				var key int
				var val *storage.Logentry
				for key, val = range entries {
					fmt.Printf("Date: %v is: %v and has the id %v\n", val.Date, val.Entry, key)
				}

			},
		},
	}

	app.Run(os.Args)

}
