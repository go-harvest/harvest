// Harvest Time Tracking CLI
//
// Credit
//
// Charney Kaye
// https://w.charney.io
//
// Outright Mental
// https://w.outright.io
//
package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "harvest"
	app.Usage = "Harvest Time Tracking CLI"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{cli.Author{Name: "Charney Kaye", Email: "hiya@charney.io"}}
	app.Commands = commands
	app.Run(os.Args)
}

var commands = []cli.Command{
	// Help
	{
		Name:    "help",
		Aliases: []string{"h", "?"},
		Usage:   "show the help",
		Action: func(c *cli.Context) {
			cli.ShowAppHelp(c)
		},
	},
	// Version
	{
		Name:    "version",
		Aliases: []string{"v", "-v"},
		Usage:   "show the version",
		Action: func(c *cli.Context) {
			cli.ShowVersion(c)
		},
	},
}
