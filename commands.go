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

var commandOpen = cli.Command{
	Name:    "open",
	Aliases: []string{"o"},
	Usage:   "Open github repository page",
	Description: `
Open URL link that can be identified by git-remote command.
If you do not specify the path, you can see the top page.
And if you specify, you can see the path on Web with highlighted lines.
`,
	Action: doOpen,
	Flags:  openFlags,
}

var openFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "root",
		Value: ".",
		Usage: "Specify root dir for repository",
	},
	cli.StringFlag{
		Name:  "ref, r",
		Value: "master",
		Usage: "Ref to open such as branch, tag and hash",
	},
	cli.IntFlag{
		Name:  "from, f",
		Usage: "Line to highlight from",
	},
	cli.IntFlag{
		Name:  "to, t",
		Usage: "Line to highlight to",
	},
	cli.BoolFlag{
		Name:  "print, p",
		Usage: "Print URL to STDOUT",
	},
}

func doOpen(c *cli.Context) {
	argPath := c.Args().Get(0)
	root := c.String("root")
	ref := c.String("ref")
	from := c.Int("from")
	to := c.Int("to")
	print := c.Bool("print")

	url, err := RemoteURL(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "remote url not found: %s\n", err)
	} else {
		urlString := url.SourceURL(ref, argPath, from, to)
		if print {
			fmt.Print(urlString)
		} else {
			fmt.Printf("opening url: \"%s\"...\n", urlString)
			cmd := exec.Command("open", urlString)
			cmd.Run()
		}
	}
}
