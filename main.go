package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/vbardakos/fython/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome to Fython REPL, %s!\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
