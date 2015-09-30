package main

import (
	"fmt"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandOpen,
}

var openFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "ref, r",
		Value: "master",
		Usage: "Ref to open such as branch, tag and hash",
	},
}

var commandOpen = cli.Command{
	Name:  "open",
	Usage: "Open github repository page",
	Description: `
    hogehoge
`,
	Action: doOpen,
	Flags:  openFlags,
}

func doOpen(c *cli.Context) {
	argDir := c.Args().Get(0)
	ref := c.String("ref")

	if argDir == "" {
		argDir = "."
	}

	url, err := NewURL(argDir, ref)
	fmt.Println(url, err)
}
