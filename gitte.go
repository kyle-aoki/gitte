package main

import (
	"os/exec"
	"strings"

	"github.com/kyle-aoki/uu"
)

func main() {
	defer uu.MainRecover()
	c("git add -A")
	c("git commit -m \"O\"")
	c("git push")
}

func c(command string) {
	commandParts := strings.Split(command, " ")
	cmd := exec.Command(commandParts[0], commandParts[1:]...)
	err := cmd.Run()
	uu.MustExec(err)
}
