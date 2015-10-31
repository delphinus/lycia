package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
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

	url, err := RemoteURL(argDir, ref)
	if err != nil {
		fmt.Fprintf(os.Stderr, "remote url not found: %s\n", err)
	} else {
		fmt.Printf("opening url: \"%s\"...\n", url.String())
		cmd := exec.Command("open", url.String())
		cmd.Run()
	}
}
