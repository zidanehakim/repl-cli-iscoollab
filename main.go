package main

import (
	"bufio"
	"fmt"
	"os"
	"repl-cli-iscoollab/cmd/commands"
	"repl-cli-iscoollab/internal/utils"
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
		// Parse input, accept extra spaces and quotes
		args := utils.ParseInput(command)

		switch args[0] {
		case "register":
			output, err := commands.Register(args[1:])
			if err != nil {
				if strings.Contains(err.Error(), "Usage: ") {
					fmt.Fprintf(os.Stderr, "%s\n", err)
				} else {
					fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				}
			}

			fmt.Print(output)
		case "create-folder":
			output, err := commands.CreateFolder(args[1:])
			if err != nil {
				if strings.Contains(err.Error(), "Usage: ") {
					fmt.Fprintf(os.Stderr, "%s\n", err)
				} else {
					fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				}
			}

			fmt.Print(output)
		case "list-folders":
			output, err := commands.ListFolders(args[1:])
			if err != nil {
				if strings.Contains(err.Error(), "Usage: ") {
					fmt.Fprintf(os.Stderr, "%s\n", err)
				} else {
					fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				}
			}

			fmt.Print(output)
		case "delete-folder":
			output, err := commands.DeleteFolder(args[1:])
			if err != nil {
				if strings.Contains(err.Error(), "Usage: ") {
					fmt.Fprintf(os.Stderr, "%s\n", err)
				} else {
					fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				}
			}

			fmt.Print(output)
		case "rename-folder":
			output, err := commands.RenameFolder(args[1:])
			if err != nil {
				if strings.Contains(err.Error(), "Usage: ") {
					fmt.Fprintf(os.Stderr, "%s\n", err)
				} else {
					fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				}
			}

			fmt.Print(output)

		case "create-file":
			output, err := commands.CreateFile(args[1:])
			if err != nil {
				if strings.Contains(err.Error(), "Usage: ") {
					fmt.Fprintf(os.Stderr, "%s\n", err)
				} else {
					fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				}
			}

			fmt.Print(output)
		case "list-files":
			output, err := commands.ListFiles(args[1:])
			if err != nil {
				if strings.Contains(err.Error(), "Usage: ") {
					fmt.Fprintf(os.Stderr, "%s\n", err)
				} else {
					fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				}
			}

			fmt.Print(output)
		case "delete-file":
			output, err := commands.DeleteFile(args[1:])
			if err != nil {
				if strings.Contains(err.Error(), "Usage: ") {
					fmt.Fprintf(os.Stderr, "%s\n", err)
				} else {
					fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				}
			}

			fmt.Print(output)
		case "help":
			output := commands.Help()
			fmt.Print(output)
		case "exit":
			commands.Exit()
		default:
			fmt.Fprintf(os.Stderr, "Error: Unrecognized command\n")
		}
	}
}
