package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

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

		splitCommand := strings.Split(command, " ")

		if splitCommand[0] == "echo" {
			echoPhrase := strings.Join(splitCommand[1:], " ")
			fmt.Fprintf(os.Stdout, "%s\n", echoPhrase)
			fmt.Fprint(os.Stdout, "$ ")
			continue
		}

		if splitCommand[0] == "type" {
			kwargs := map[string]string{
				"help": "help",
				"exit": "exit",
				"echo": "echo",
				"type": "type",
			}

			val, ok := kwargs[splitCommand[1]]
			if ok {
				fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", val)
				fmt.Fprint(os.Stdout, "$ ")
				continue
			} else {
				fmt.Fprintf(os.Stdout, "%s: not found\n", splitCommand[1])
				fmt.Fprint(os.Stdout, "$ ")
				continue
			}

		}

		if command == "exit 0" {
			os.Exit(0)
		}

		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		fmt.Fprint(os.Stdout, "$ ")

	}

}
