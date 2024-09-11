package user

import (
	"fmt"
	"repl-cli-iscoollab/internal/utils"
	"sort"
	"time"
)

var (
	ListUser = make(map[string]*User)

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

const (
	MaxUsernameLength   = 25
	MaxFolderNameLength = 255
	MaxFileNameLength   = 255
)

type User struct {
	Username string
	Folders  map[string]*Folder
}

func (u *User) CreateFolder(folderName string, description string) error {
	if !utils.ValidateString(folderName) {
		return fmt.Errorf("%s contain invalid chars", folderName)
	}

	if _, exists := u.Folders[folderName]; exists {
		return fmt.Errorf("the %s has already existed", folderName)
	}

	folder := &Folder{
		Name:        folderName,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		Description: description,
		Files:       make(map[string]*File),
	}

	u.Folders[folderName] = folder

	return nil
}

func (u *User) DeleteFolder(folderName string) error {
	if _, exists := u.Folders[folderName]; exists {
		delete(u.Folders, folderName)
		return nil
	}

	return fmt.Errorf("the %s doesn't exist", folderName)
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

	folders := make([]*Folder, 0, len(u.Folders))
	for _, folder := range u.Folders {
		folders = append(folders, folder)
	}

	switch sortBy {
	case "--sort-name":
		sort.Slice(folders, func(i, j int) bool {
			if isAsc {
				return folders[i].Name < folders[j].Name
			}
			return folders[i].Name > folders[j].Name
		})
	case "--sort-created":
		sort.Slice(folders, func(i, j int) bool {
			if isAsc {
				return folders[i].CreatedAt < folders[j].CreatedAt
			}
			return folders[i].CreatedAt > folders[j].CreatedAt
		})
	default:
		return nil, fmt.Errorf(CommandsUsage["list-folders"])
	}

	return folders, nil
}

func (u *User) RenameFolder(folderName string, newFolderName string) error {
	folder, exists := u.Folders[folderName]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", folderName)
	}

	if _, exists := u.Folders[newFolderName]; exists {
		return fmt.Errorf("the %s already exists", newFolderName)
	}

	folder.Name = newFolderName
	u.Folders[newFolderName] = folder
	delete(u.Folders, folderName)

	return nil
}

func (u *User) GetFolder(folderName string) (*Folder, error) {
	folder, exists := u.Folders[folderName]
	if !exists {
		return nil, fmt.Errorf("the %s doesn't exist", folderName)
	}
	return folder, nil
}

func RegisterUser(username string) error {
	if _, exists := ListUser[username]; exists {
		return fmt.Errorf("the %s has already existed", username)
	}

	if !utils.ValidateString(username) {
		return fmt.Errorf("the %s contain invalid chars", username)
	}

	newUser := &User{
		Username: username,
		Folders:  make(map[string]*Folder),
	}

	ListUser[username] = newUser

	return nil
}

func GetUser(username string) (*User, error) {
	if user, exists := ListUser[username]; exists {
		return user, nil
	}

	return nil, fmt.Errorf("the %s doesn't exist", username)
}
