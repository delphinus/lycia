package main

import (
	"github.com/codegangsta/cli"
	"os"
)

var Version = "HEAD"

func main() {
	newApp().Run(os.Args)
}

func newApp() (app *cli.App) {
	app = cli.NewApp()
	app.Name = "lycia"
	app.Usage = "Open Github repository page"
	app.Version = Version
	app.Author = "delphinus"
	app.Email = "delphinus@remora.cx"
	app.Commands = Commands
	return
}
