package main

import (
	"bytes"
	"io"
	"os"
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
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)
	commandParts := strings.Split(command, " ")
	cmd := exec.Command(commandParts[0], commandParts[1:]...)
	cmd.Stdout = mw
	cmd.Stderr = mw
	err := cmd.Run()
	uu.MustExec(err)
}
