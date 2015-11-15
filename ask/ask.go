package ask

import (
	"bufio"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

var normal = color.New(color.FgGreen).PrintfFunc()

func Ask(question string, toHide bool) (answer string, err error) {
	normal(question + " ")
	if toHide {
		// TODO: This function makes no sense in non-UNIX platform.
		byt, err := terminal.ReadPassword(0)
		if err == nil {
			answer = string(byt)
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		answer, err = reader.ReadString('\n')
		answer = strings.TrimRight(answer, "\n")
	}
	return
}
