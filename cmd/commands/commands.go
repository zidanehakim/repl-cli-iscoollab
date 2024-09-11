package commands

import (
	"flag"
	"fmt"
	"os"
	"repl-cli-iscoollab/internal/user"
)

func Register(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf(user.CommandsUsage["register"])
	}

	username := args[0]
	err := user.RegisterUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("Add %s successfully.\n", username)

	return nil
}

func CreateFolder(args []string) error {
	if len(args) != 2 && len(args) != 3 {
		return fmt.Errorf(user.CommandsUsage["create-folder"])
	}

	username := args[0]
	folderName := args[1]
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

	fmt.Printf("Create %s successfully.\n", folderName)

	return nil
}

func DeleteFolder(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf(user.CommandsUsage["delete-folder"])
	}

	username := args[0]
	folderName := args[1]

	user, err := user.GetUser(username)
	if err != nil {
		return err
	}

	err = user.DeleteFolder(folderName)
	if err != nil {
		return err
	}

	fmt.Printf("Delete %s successfully.\n", folderName)

	return nil
}

func ListFolders(args []string) error {
	if len(args) < 1 || len(args) > 3 {
		return fmt.Errorf(user.CommandsUsage["list-folders"])
	}
	if args[1] != "--sort-name" && args[1] != "--sort-created" {
		return fmt.Errorf(user.CommandsUsage["list-files"])
	}
	if args[2] != "asc" && args[2] != "desc" {
		return fmt.Errorf(user.CommandsUsage["list-files"])
	}

	username := args[0]
	sortBy := "--sort-name"
	sortOrder := "asc"
	if len(args) > 1 {
		sortBy = args[1]
	}
	if len(args) > 2 {
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
		if folder.Description != "" {
			folder.Description = " " + folder.Description
		}
		fmt.Printf("%s%s %s %s\n", folder.Name, folder.Description, folder.CreatedAt, user.Username)
	}

	return nil
}

func RenameFolder(args []string) error {
	if len(args) != 3 {
		return fmt.Errorf(user.CommandsUsage["rename-folder"])
	}

	username := args[0]
	folderName := args[1]
	newFolderName := args[2]

	user, err := user.GetUser(username)
	if err != nil {
		return err
	}

	err = user.RenameFolder(folderName, newFolderName)
	if err != nil {
		return err
	}

	fmt.Printf("Rename %s to %s successfully.\n", folderName, newFolderName)

	return nil
}

func CreateFile(args []string) error {
	if len(args) != 3 && len(args) != 4 {
		return fmt.Errorf(user.CommandsUsage["create-file"])
	}

	username := args[0]
	folderName := args[1]
	fileName := args[2]
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

	fmt.Printf("Create %s in %s/%s successfully.\n", fileName, username, folderName)

	return nil
}

func ListFiles(args []string) error {
	if len(args) < 2 || len(args) > 4 {
		return fmt.Errorf(user.CommandsUsage["list-files"])
	}
	if args[1] != "--sort-name" && args[1] != "--sort-created" {
		return fmt.Errorf(user.CommandsUsage["list-files"])
	}
	if args[2] != "asc" && args[2] != "desc" {
		return fmt.Errorf(user.CommandsUsage["list-files"])
	}

	username := args[0]
	folderName := args[1]
	sortBy := "--sort-name"
	sortOrder := "asc"
	if len(args) > 1 {
		sortBy = args[1]
	}
	if len(args) > 2 {
		sortOrder = args[2]
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
		if file.Description != "" {
			file.Description = " " + file.Description
		}
		fmt.Printf("%s%s %s %s\n", file.Name, file.Description, file.CreatedAt, user.Username)
	}

	return nil
}

func DeleteFile(args []string) error {
	if len(args) != 3 {
		return fmt.Errorf(user.CommandsUsage["delete-file"])
	}

	username := args[0]
	folderName := args[1]
	fileName := args[2]

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

	fmt.Printf("Deleted file %s from %s/%s successfully.\n", fileName, username, folderName)

	return nil
}

func Help() {
	fmt.Println("Help")
}

func Exit() {
	fmt.Print("\033[H\033[2J")
	os.Exit(0)
}
