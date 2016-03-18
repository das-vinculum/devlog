package main

import (
	"github.com/codegangsta/cli"
	"os"
	"strings"
	"time"
)

type logentry struct {
	entry string
	date  time.Time
}

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
				homeDir := os.Getenv("HOME")
				logFile := homeDir + "/" + "done.log.txt"
				f, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
				check(err)
				defer f.Close()
				f.WriteString(time.Now().Format(time.UnixDate))
				f.WriteString("|")
				f.WriteString(strings.Join(c.Args()[:], " "))
				f.WriteString("\n")
				f.Sync()
				println("Added ", strings.Join(c.Args()[:], " "))
				println("Recorded for", time.Now().Format(time.UnixDate))

			},
		},
	}

	app.Run(os.Args)

}
