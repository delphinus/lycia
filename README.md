# lycia [![Build Status](https://travis-ci.org/delphinus35/lycia.svg?branch=master)](https://travis-ci.org/delphinus35/lycia)

```
NAME:
   lycia - Open Github repository page

USAGE:
   lycia [global options] command [command options]

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
     - open PR #50

VERSION:
   v0.0.4

AUTHOR(S):
   delphinus <delphinus@remora.cx>

COMMANDS:
   open, o              Open github repository page
   issue, i             Open github issue page
   pullrequest, p, pull Open github pullrequest page
   help, h              Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h           show help
   --version, -v        print the version
```
