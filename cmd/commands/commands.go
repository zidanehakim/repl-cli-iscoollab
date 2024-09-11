package commands

import (
	"flag"
	"fmt"
	"os"
	"repl-cli-iscoollab/internal/user"
	"strings"
)

func Register(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf(user.CommandsUsage["register"])
	}

	username := strings.ToLower(args[0])
	err := user.RegisterUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("Add %s successfully\n", username)

	return nil
}

func CreateFolder(args []string) error {
	if len(args) != 2 && len(args) != 3 {
		return fmt.Errorf(user.CommandsUsage["create-folder"])
	}

	username := strings.ToLower(args[0])
	folderName := strings.ToLower(args[1])
	var description string
	if len(args) > 2 {
		description = args[2]
	}
	flag.Parse()

	user, err := user.GetUser(username)
	if err != nil {
		return err
	}

	err = user.CreateFolder(folderName, description)
	if err != nil {
		return err
	}

	fmt.Printf("Create %s successfully\n", folderName)

	return nil
}

func DeleteFolder(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf(user.CommandsUsage["delete-folder"])
	}

	username := strings.ToLower(args[0])
	folderName := strings.ToLower(args[1])

	user, err := user.GetUser(username)
	if err != nil {
		return err
	}

	err = user.DeleteFolder(folderName)
	if err != nil {
		return err
	}

	fmt.Printf("Delete %s successfully\n", folderName)

	return nil
}

func ListFolders(args []string) error {
	if len(args) < 1 || len(args) > 3 {
		return fmt.Errorf(user.CommandsUsage["list-folders"])
	}

	username := strings.ToLower(args[0])
	sortBy := "--sort-name"
	sortOrder := "asc"
	if len(args) > 1 {
		if args[1] != "--sort-name" && args[1] != "--sort-created" {
			return fmt.Errorf(user.CommandsUsage["list-folders"])
		}

		sortBy = args[1]
	}
	if len(args) > 2 {
		if args[2] != "asc" && args[2] != "desc" {
			return fmt.Errorf(user.CommandsUsage["list-folders"])
		}

		sortOrder = args[2]
	}

	user, err := user.GetUser(username)
	if err != nil {
		return err
	}

	folders, err := user.ListFolders(sortBy, sortOrder)
	if err != nil {
		return err
	}

	for _, folder := range folders {
		var description string
		if folder.Description != "" {
			description = " " + folder.Description
		}
		fmt.Printf("%s%s %s %s\n", folder.Name, description, folder.CreatedAt, user.Username)
	}

	return nil
}

func RenameFolder(args []string) error {
	if len(args) != 3 {
		return fmt.Errorf(user.CommandsUsage["rename-folder"])
	}

	username := strings.ToLower(args[0])
	folderName := strings.ToLower(args[1])
	newFolderName := args[2]

	user, err := user.GetUser(username)
	if err != nil {
		return err
	}

	err = user.RenameFolder(folderName, newFolderName)
	if err != nil {
		return err
	}

	fmt.Printf("Rename %s to %s successfully\n", folderName, newFolderName)

	return nil
}

func CreateFile(args []string) error {
	if len(args) != 3 && len(args) != 4 {
		return fmt.Errorf(user.CommandsUsage["create-file"])
	}

	username := strings.ToLower(args[0])
	folderName := strings.ToLower(args[1])
	fileName := strings.ToLower(args[2])
	var description string
	if len(args) > 3 {
		description = args[3]
	}

	user, err := user.GetUser(username)
	if err != nil {
		return err
	}

	folder, err := user.GetFolder(folderName)
	if err != nil {
		return err
	}

	err = folder.CreateFile(fileName, description)
	if err != nil {
		return err
	}

	fmt.Printf("Create %s in %s/%s successfully\n", fileName, username, folderName)

	return nil
}

func ListFiles(args []string) error {
	if len(args) < 2 || len(args) > 4 {
		return fmt.Errorf(user.CommandsUsage["list-files"])
	}

	username := strings.ToLower(args[0])
	folderName := strings.ToLower(args[1])
	sortBy := "--sort-name"
	sortOrder := "asc"
	if len(args) > 2 {
		if args[2] != "--sort-name" && args[2] != "--sort-created" {
			return fmt.Errorf(user.CommandsUsage["list-files"])
		}

		sortBy = args[2]
	}
	if len(args) > 3 {
		if args[3] != "asc" && args[3] != "desc" {
			return fmt.Errorf(user.CommandsUsage["list-files"])
		}

		sortOrder = args[3]
	}

	user, err := user.GetUser(username)
	if err != nil {
		return err
	}

	folder, err := user.GetFolder(folderName)
	if err != nil {
		return err
	}

	files, err := folder.ListFiles(sortBy, sortOrder)
	if err != nil {
		return err
	}

	for _, file := range files {
		var description string
		if file.Description != "" {
			description = " " + file.Description
		}
		fmt.Printf("%s%s %s %s\n", file.Name, description, file.CreatedAt, user.Username)
	}

	return nil
}

func DeleteFile(args []string) error {
	if len(args) != 3 {
		return fmt.Errorf(user.CommandsUsage["delete-file"])
	}

	username := strings.ToLower(args[0])
	folderName := strings.ToLower(args[1])
	fileName := strings.ToLower(args[2])

	user, err := user.GetUser(username)
	if err != nil {
		return err
	}

	folder, err := user.GetFolder(folderName)
	if err != nil {
		return err
	}

	err = folder.DeleteFile(fileName)
	if err != nil {
		return err
	}

	fmt.Printf("Deleted file %s from %s/%s successfully\n", fileName, username, folderName)

	return nil
}

func Help() {
	fmt.Println("Available commands:")
	fmt.Println("  register [username]                                                         - Register a new user")
	fmt.Println("  create-folder [username] [foldername] [description]?                        - Create a new folder")
	fmt.Println("  list-folders [username] [--sort-name|--sort-created] [asc|desc]             - List folders for a user")
	fmt.Println("  delete-folder [username] [foldername]                                       - Delete a folder")
	fmt.Println("  rename-folder [username] [foldername] [new-folder-name]                     - Rename a folder")
	fmt.Println("  create-file [username] [foldername] [filename] [description]?               - Create a new file")
	fmt.Println("  list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]  - List files in a folder")
	fmt.Println("  delete-file [username] [foldername] [filename]                              - Delete a file")
	fmt.Println("  help                                                                        - Show this help message")
	fmt.Println("  exit                                                                        - Exit the program")
	fmt.Println("\nNote: Parameters in square brackets [] are required, those with ? are optional.")
	fmt.Println("For sorting, you can use either --sort-name or --sort-created, followed by asc (ascending) or desc (descending).")
}

func Exit() {
	fmt.Print("\033[H\033[2J")
	os.Exit(0)
}
