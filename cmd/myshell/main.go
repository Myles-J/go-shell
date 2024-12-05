package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Uncomment this block to pass the first stage

	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	f := bufio.NewReader(os.Stdin)

	for {
		command, err := f.ReadString('\n')
		command = strings.Trim(command, "\n")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		fmt.Fprint(os.Stdout, "$ ")
	}

}