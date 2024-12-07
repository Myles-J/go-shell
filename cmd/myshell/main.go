package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type (
	Args    []string
	CmdFunc func(Args)
)

var commands = make(map[string]CmdFunc)

func registerCommand(cmd string, fn CmdFunc) {
	commands[cmd] = fn
}

func exit(args Args) {
	if len(args) == 0 {
		os.Exit(1)
	}
	if code, err := strconv.Atoi(args[0]); err == nil {
		os.Exit(code)
	}
}

func returnCommandFilePath(command string) (string, error) {
	paths := strings.Split(os.Getenv("PATH"), ":")

	for _, path := range paths {
		fp := filepath.Join(path, command)
		if _, err := os.Stat(fp); err == nil {
			return fp, nil
		}
	}
	return "", errors.New("command not found")
}

func typer(args Args) {
	if len(args) == 0 {
		fmt.Println("")
	}

	_, builtin := commands[args[0]]

	if builtin {
		fmt.Printf("%s is a shell builtin\n", args[0])
		return
	}

	path, err := returnCommandFilePath(args[0])

	if err != nil {
		fmt.Printf("%s: not found\n", args[0])
		return
	}

	fmt.Println(path)
}

func echo(args Args) {
	phrase := strings.Join(args, " ")
	fmt.Fprintf(os.Stdout, "%s\n", phrase)
}

func pwd(args Args) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("error getting pwd: %s", err.Error())
		os.Exit(1)
	}
	fmt.Println(pwd)
}

func notFound(cmd string) {
	fmt.Printf("%s: command not found\n", cmd)
}

func initCommands() {
	registerCommand("echo", echo)
	registerCommand("exit", exit)
	registerCommand("type", typer)
	registerCommand("pwd", pwd)
}

func main() {
	initCommands()
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		in, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Printf("error reading from stdin: %s", err.Error())
			os.Exit(1)
		}

		inputs := strings.Split(strings.TrimSpace(in), " ")
		cmd := inputs[0]
		args := inputs[1:]

		path, err := returnCommandFilePath(cmd)
		if err == nil {
			// Add the command to the command list
			commands[cmd] = func(args Args) {
				cmd := exec.Command(path, args...)
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
			}
		}

		cmdFn, ok := commands[cmd]

		if !ok {
			notFound(cmd)
		} else {
			cmdFn(args)
		}
	}
}
