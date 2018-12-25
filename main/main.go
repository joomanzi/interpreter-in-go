package main

import (
	"fmt"
	"github.com/interpreter-in-go/repl"
	"os"
	user2 "os/user"
)

func main(){
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! \n", user.Username)
	fmt.Printf("type in command\n")
	repl.Start(os.Stdin, os.Stdout)
}
