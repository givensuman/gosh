package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var history []string = make([]string, 0)

	for {
		workingDirectory, _ := os.Getwd()
		hostName, _ := os.Hostname()
		user, _ := os.LookupEnv("USER")

		fmt.Printf("\n[%s@%s@%s]\n", user, hostName, workingDirectory)
		fmt.Print("> ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		command, err := execute(input)
		history = append(history, command)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execute(command string) (string, error) {
	input := strings.TrimSuffix(command, "\n")
	args := strings.Split(input, " ")

	switch args[0] {
	case "exit":
		os.Exit(0)

	case "cd":
		if len(args) < 2 {
			os.Chdir("/")
			return input, nil
		}
		err := os.Chdir(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	default:
		cmd := exec.Command(args[0], args[1:]...)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return input, cmd.Run()
	}

	return input, nil
}
