package commands

import (
	"flag"
	"fmt"
	"os"
	"repl-cli-iscoollab/internal/user"
	"strings"
)

func Register(args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf(user.CommandsUsage["register"])
	}

	username := strings.ToLower(args[0])
	err := user.RegisterUser(username)
	if err != nil {
		return "", err
	}

	output := fmt.Sprintf("Add %s successfully\n", username)
	return output, nil
}

func CreateFolder(args []string) (string, error) {
	if len(args) != 2 && len(args) != 3 {
		return "", fmt.Errorf(user.CommandsUsage["create-folder"])
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
		return "", err
	}

	err = user.CreateFolder(folderName, description)
	if err != nil {
		return "", err
	}

	output := fmt.Sprintf("Create %s successfully\n", folderName)
	return output, nil
}

func DeleteFolder(args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf(user.CommandsUsage["delete-folder"])
	}

	username := strings.ToLower(args[0])
	folderName := strings.ToLower(args[1])

	user, err := user.GetUser(username)
	if err != nil {
		return "", err
	}

	err = user.DeleteFolder(folderName)
	if err != nil {
		return "", err
	}

	output := fmt.Sprintf("Delete %s successfully\n", folderName)
	return output, nil
}

func ListFolders(args []string) (string, error) {
	if len(args) < 1 || len(args) > 3 {
		return "", fmt.Errorf(user.CommandsUsage["list-folders"])
	}

	username := strings.ToLower(args[0])
	sortBy := "--sort-name"
	sortOrder := "asc"
	if len(args) > 1 {
		if args[1] != "--sort-name" && args[1] != "--sort-created" {
			return "", fmt.Errorf(user.CommandsUsage["list-folders"])
		}

		sortBy = args[1]
	}
	if len(args) > 2 {
		if args[2] != "asc" && args[2] != "desc" {
			return "", fmt.Errorf(user.CommandsUsage["list-folders"])
		}

		sortOrder = args[2]
	}

	user, err := user.GetUser(username)
	if err != nil {
		return "", err
	}

	folders, err := user.ListFolders(sortBy, sortOrder)
	if err != nil {
		return "", err
	}

	var output strings.Builder
	for _, folder := range folders {
		var description string
		if folder.Description != "" {
			description = " " + folder.Description
		}
		output.WriteString(fmt.Sprintf("%s%s %s %s\n", folder.Name, description, folder.CreatedAt, user.Username))
	}

	return output.String(), nil
}

func RenameFolder(args []string) (string, error) {
	if len(args) != 3 {
		return "", fmt.Errorf(user.CommandsUsage["rename-folder"])
	}

	username := strings.ToLower(args[0])
	folderName := strings.ToLower(args[1])
	newFolderName := strings.ToLower(args[2])

	user, err := user.GetUser(username)
	if err != nil {
		return "", err
	}

	err = user.RenameFolder(folderName, newFolderName)
	if err != nil {
		return "", err
	}

	output := fmt.Sprintf("Rename %s to %s successfully\n", folderName, newFolderName)
	return output, nil
}

func CreateFile(args []string) (string, error) {
	if len(args) != 3 && len(args) != 4 {
		return "", fmt.Errorf(user.CommandsUsage["create-file"])
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
		return "", err
	}

	folder, err := user.GetFolder(folderName)
	if err != nil {
		return "", err
	}

	err = folder.CreateFile(fileName, description)
	if err != nil {
		return "", err
	}

	output := fmt.Sprintf("Create %s in %s/%s successfully\n", fileName, username, folderName)
	return output, nil
}

func ListFiles(args []string) (string, error) {
	if len(args) < 2 || len(args) > 4 {
		return "", fmt.Errorf(user.CommandsUsage["list-files"])
	}

	username := strings.ToLower(args[0])
	folderName := strings.ToLower(args[1])
	sortBy := "--sort-name"
	sortOrder := "asc"
	if len(args) > 2 {
		if args[2] != "--sort-name" && args[2] != "--sort-created" {
			return "", fmt.Errorf(user.CommandsUsage["list-files"])
		}

		sortBy = args[2]
	}
	if len(args) > 3 {
		if args[3] != "asc" && args[3] != "desc" {
			return "", fmt.Errorf(user.CommandsUsage["list-files"])
		}

		sortOrder = args[3]
	}

	user, err := user.GetUser(username)
	if err != nil {
		return "", err
	}

	folder, err := user.GetFolder(folderName)
	if err != nil {
		return "", err
	}

	files, err := folder.ListFiles(sortBy, sortOrder)
	if err != nil {
		return "", err
	}

	var output strings.Builder
	for _, file := range files {
		var description string
		if file.Description != "" {
			description = " " + file.Description
		}
		output.WriteString(fmt.Sprintf("%s%s %s %s\n", file.Name, description, file.CreatedAt, user.Username))
	}

	return output.String(), nil
}

func DeleteFile(args []string) (string, error) {
	if len(args) != 3 {
		return "", fmt.Errorf(user.CommandsUsage["delete-file"])
	}

	username := strings.ToLower(args[0])
	folderName := strings.ToLower(args[1])
	fileName := strings.ToLower(args[2])

	user, err := user.GetUser(username)
	if err != nil {
		return "", err
	}

	folder, err := user.GetFolder(folderName)
	if err != nil {
		return "", err
	}

	err = folder.DeleteFile(fileName)
	if err != nil {
		return "", err
	}

	output := fmt.Sprintf("Deleted file %s from %s/%s successfully\n", fileName, username, folderName)
	return output, nil
}

func Help() string {
	output := `Available commands:
  register [username]                                                         - Register a new user
  create-folder [username] [foldername] [description]?                        - Create a new folder
  list-folders [username] [--sort-name|--sort-created] [asc|desc]             - List folders for a user
  delete-folder [username] [foldername]                                       - Delete a folder
  rename-folder [username] [foldername] [new-folder-name]                     - Rename a folder
  create-file [username] [foldername] [filename] [description]?               - Create a new file
  list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]  - List files in a folder
  delete-file [username] [foldername] [filename]                              - Delete a file
  help                                                                        - Show this help message
  exit                                                                        - Exit the program

Note: Parameters in square brackets [] are required, those with ? are optional.
For sorting, you can use either --sort-name or --sort-created, followed by asc (ascending) or desc (descending).
`
	return output
}

func Exit() {
	fmt.Print("\033[H\033[2J")
	os.Exit(0)
}
