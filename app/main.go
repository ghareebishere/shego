package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
	fmt.Print("$ ")
	// so we need to use read string to get the commands from the user
	// so lets declare a reader
	reader := bufio.NewReader(os.Stdin)

	command, err := reader.ReadString('\n')
	if err  != nil {
		// fmt.Println("input is wrong or couldn't be read")
		// using stderr here and passing the error with exit code would be useful
		// we also used the Fprintf to fromat the error
		fmt.Fprintf(os.Stderr, "%v error while reading the command: ", err)
	}
	// remove the trailing \n out from the command
	fmt.Println(command[:len(command)-1] + " command not found")
}
