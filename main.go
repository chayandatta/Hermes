package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/chayandatta/Hermes/repl"
)

func main() {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! welcome to the repl of Hermes\n", user.Username)
	fmt.Print("Start typing to see the repl in action\n")
	repl.Start(os.Stdin, os.Stdout)
}
