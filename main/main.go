package main

import (
	"fmt"
	"os"
	"os/user"
	"gordon/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! I am GOrdon!\n", user.Username)
	fmt.Printf("Type in some commands here\n")
	repl.Start(os.Stdin, os.Stdout)
}