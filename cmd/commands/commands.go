package commands

import (
	"fmt"
	"os"
	"repl-cli-iscoollab/internal/user"
)

func Register(username string) error {
	err := user.RegisterUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("Add %s successfully.\n", username)

	return nil
}

func CreateFolder(username string, folderName string, description string) error {
	user, err := user.GetUser(username)
	if err != nil {
		return err
	}

	user.CreateFolder(folderName, description)

	fmt.Printf("Create %s successfully.\n", folderName)

	return nil
}

func DeleteFolder(username string, folderName string) error {
	user, err := user.GetUser(username)
	if err != nil {
		return err
	}

	user.DeleteFolder(folderName)

	fmt.Printf("Delete %s successfully.\n", folderName)

	return nil
}

func ListFolders(username string, sortBy string, sortOrder string) error {
	user, err := user.GetUser(username)
	if err != nil {
		return err
	}

	folders, err := user.ListFolders(sortBy, sortOrder)
	if err != nil {
		return err
	}

	for _, folder := range folders {
		fmt.Printf("%s %s %s %s\n", folder.Name, folder.Description, folder.CreatedAt, user.Username)
	}

	return nil
}

func Help() {
	fmt.Println("Help")
}

func Exit() {
	fmt.Print("\033[H\033[2J")
	os.Exit(0)
}
