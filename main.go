package main

import (
	"github.com/codegangsta/cli"
	"os"
)

// Version is a version number
var Version = "v0.0.7"

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
	app.ArgsUsage = `

   # examples
   lycia open
     - show project top page if you are on top of project directory

   lycia open --print
     - show URL to STDOUT instead of opening in browser

   lycia open --root /path/to/repository
     - you can specify the path of repository

   lycia open relative/path/to/source.go
     - open source.go of master branch on github

   lycia open relative/path/to/source.go --ref develop --from 30 --to 32
     - open source.go of develop branch on github with highlighted lines 30 to 32

   lycia o relative/path/to/source.go -r develop -f 30 -t 32
     - same in short form

   lycia issue 40
     - open issue #40

   lycia pullrequest #50
     - open PR #50`
	return
}
