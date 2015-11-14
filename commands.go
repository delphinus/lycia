package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/delphinus35/lycia/fssh"
	"github.com/delphinus35/lycia/github"
	"os"
	"os/exec"
	"strconv"
)

func openOrPrintURL(c *cli.Context, urlString string, doPrint bool) {
	if doPrint {
		fmt.Print(urlString)
	} else {
		fmt.Printf("opening url: \"%s\"...\n", urlString)
		noFssh := c.Bool("no-fssh")
		var cmd *exec.Cmd
		if noFssh {
			cmd = exec.Command("open", urlString)
		} else {
			cmd = fssh.Command("open", urlString)
		}
		cmd.Run()
	}
}

var Commands = []cli.Command{
	commandOpen,
	commandIssue,
	commandPullrequest,
}

var commonFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "root",
		Value: ".",
		Usage: "Specify root dir for repository",
	},
	cli.BoolFlag{
		Name:  "print, p",
		Usage: "Print URL to STDOUT",
	},
	cli.BoolFlag{
		Name:  "no-fssh",
		Usage: "Do not use fssh execution",
	},
}

var commandOpen = cli.Command{
	Name:    "open",
	Aliases: []string{"o"},
	Usage:   "Open github repository page",
	Description: `Open URL link that can be identified by git-remote command.
   If you do not specify the path, you can see the top page.
   And if you specify, you can see the path on Web with highlighted lines.`,
	Action: doOpen,
	Flags:  openFlags,
}

var openFlags = append(
	commonFlags,
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
)

func doOpen(c *cli.Context) {
	argPath := c.Args().Get(0)
	root := c.String("root")
	ref := c.String("ref")
	from := c.Int("from")
	to := c.Int("to")
	doPrint := c.Bool("print")

	remoteURL, err := RemoteURL(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "remote url not found: %s\n", err)
	} else {
		urlString := remoteURL.SourceURL(ref, argPath, from, to)
		openOrPrintURL(c, urlString, doPrint)
	}
}

var commandIssue = cli.Command{
	Name:    "issue",
	Aliases: []string{"i"},
	Usage:   "Open github issue page",
	Description: `Open issue page that is specified by number in Args.
   If number is not specified, it will open the top page of issues.`,
	Action: doIssue,
	Flags:  issueFlags,
}

var issueFlags = commonFlags

func doIssue(c *cli.Context) {
	argNumber, _ := strconv.Atoi(c.Args().Get(0))
	root := c.String("root")
	doPrint := c.Bool("print")

	remoteURL, err := RemoteURL(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "remote url not found: %s\n", err)
	} else {
		urlString := remoteURL.IssueURL(argNumber)
		openOrPrintURL(c, urlString, doPrint)
	}
}

var commandPullrequest = cli.Command{
	Name:    "pullrequest",
	Aliases: []string{"p", "pr", "pull"},
	Usage:   "Open github pullrequest page",
	Description: `Open pullrequest page that is specified by number in Args.
   If number is not specified, it will open the top page of pullrequests.`,
	Action: doPullrequest,
	Flags:  pullrequestFlags,
}

var pullrequestFlags = append(commonFlags,
	cli.StringFlag{
		Name:  "branch, b",
		Value: "",
		Usage: "Specify head branch for pullrequest to open",
	},
)

func doPullrequest(c *cli.Context) {
	argNumber, _ := strconv.Atoi(c.Args().Get(0))
	root := c.String("root")
	doPrint := c.Bool("print")
	branch := c.String("branch")

	remoteURL, err := RemoteURL(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "remote url not found: %s\n", err)
		return
	}

	if argNumber != 0 {
		urlString := remoteURL.PullrequestURL(argNumber)
		openOrPrintURL(c, urlString, doPrint)
		return
	}

	repo := github.Repository(remoteURL.ToURL())
	prURL, err := repo.PullrequestURL(branch)
	if err != nil {
		fmt.Fprintf(os.Stderr, "pullrequest URL not found: %s\n", err)
	} else {
		openOrPrintURL(c, prURL.String(), doPrint)
	}
}
