package user

import (
	"fmt"
	"repl-cli-iscoollab/internal/utils"
	"sort"
	"time"
)

var (
	ListUser []*User

	CommandsUsage = map[string]string{
		"list-files":    "Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]",
		"create-file":   "Usage: create-file [username] [foldername] [filename] [description]?",
		"delete-file":   "Usage: delete-file [username] [foldername] [filename]",
		"register":      "Usage: register [username]",
		"create-folder": "Usage: create-folder [username] [foldername] [description]?",
		"list-folders":  "Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]",
		"delete-folder": "Usage: delete-folder [username] [foldername]",
		"rename-folder": "Usage: rename-folder [username] [foldername] [new-folder-name]",
		"help":          "Usage: help",
		"exit":          "Usage: exit",
	}
)

type User struct {
	Username string
	Folders  []*Folder
}

func (u *User) CreateFolder(folderName string, description string) error {
	if !utils.ValidateString(folderName) {
		return fmt.Errorf("Error: %s contain invalid chars.", folderName)
	}

	if u.checkFolder(folderName) {
		return fmt.Errorf("Error: The %s has already existed.", folderName)
	}

	folder := &Folder{
		Name:        folderName,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		Description: description,
	}

	u.Folders = append(u.Folders, folder)

	return nil
}

func (u *User) DeleteFolder(folderName string) error {
	for i, folder := range u.Folders {
		if folder.Name == folderName {
			u.Folders = append(u.Folders[:i], u.Folders[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Error: The %s doesn't exist.", folderName)
}

func (u *User) ListFolders(sortBy string, sortOrder string) ([]*Folder, error) {
	var isAsc bool
	switch sortOrder {
	case "asc":
		isAsc = true
	case "desc":
		isAsc = false
	default:
		return nil, fmt.Errorf(CommandsUsage["list-folders"])
	}

	switch sortBy {
	case "--sort-name":
		sort.Slice(u.Folders, func(i, j int) bool {
			if isAsc {
				return u.Folders[i].Name < u.Folders[j].Name
			}
			return u.Folders[i].Name > u.Folders[j].Name
		})
	case "--sort-created":
		sort.Slice(u.Folders, func(i, j int) bool {
			if isAsc {
				return u.Folders[i].CreatedAt < u.Folders[j].CreatedAt
			}
			return u.Folders[i].CreatedAt > u.Folders[j].CreatedAt
		})
	default:
		return nil, fmt.Errorf(CommandsUsage["list-folders"])
	}

	return u.Folders, nil
}

func (u *User) RenameFolder(folderName string, newFolderName string) error {
	for _, folder := range u.Folders {
		if folder.Name == folderName {
			folder.Name = newFolderName
		}
	}

	return fmt.Errorf("Error: The %s doesn't exist.", folderName)
}

func (u *User) GetFolder(folderName string) (*Folder, error) {
	for _, folder := range u.Folders {
		if folder.Name == folderName {
			return folder, nil
		}
	}
	return nil, fmt.Errorf("Error: The %s doesn't exist.", folderName)
}

func (u *User) checkFolder(folderName string) bool {
	for _, folder := range u.Folders {
		if folder.Name == folderName {
			return true
		}
	}

	return false
}

func RegisterUser(username string) error {
	if CheckUser(username) {
		return fmt.Errorf("Error: The %s has already existed.", username)
	}

	if !utils.ValidateString(username) {
		return fmt.Errorf("Error: The %s contain invalid chars.", username)
	}

	newUser := &User{
		Username: username,
		Folders:  make([]*Folder, 0),
	}

	ListUser = append(ListUser, newUser)

	return nil
}

func GetUser(username string) (*User, error) {
	for _, user := range ListUser {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, fmt.Errorf("Error: The %s doesn't exist.", username)
}

func CheckUser(username string) bool {
	for _, user := range ListUser {
		if user.Username == username {
			return true
		}
	}

	return false
}
