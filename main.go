package main

import (
	"bufio"
	"fmt"
	"os"
	"repl-cli-iscoollab/cmd/commands"
	"strings"
)

func main() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Welcome to Virtual File System Management REPL")
	fmt.Println("Type 'help' to see the list of commands")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\n> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			continue
		}

		command = strings.TrimSpace(command)
		args := strings.Split(command, " ")

		switch args[0] {
		case "register":
			username := args[1]
			err := commands.Register(username)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
		case "create-folder":
			username := args[1]
			folderName := args[2]
			description := args[3]
			err := commands.CreateFolder(username, folderName, description)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
		case "list-folders":
			username := args[1]
			sortBy := args[2]
			sortOrder := args[3]
			err := commands.ListFolders(username, sortBy, sortOrder)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
		case "delete-folder":
			username := args[1]
			folderName := args[2]
			err := commands.DeleteFolder(username, folderName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
		case "help":
			commands.Help()
		case "exit":
			commands.Exit()
		default:
			fmt.Fprintf(os.Stderr, "Error: Unrecognized command\n")
		}
	}
}
