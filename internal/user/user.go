package user

import (
	"fmt"
	"repl-cli-iscoollab/internal/utils"
	"sort"
	"time"
)

var (
	ListUser []*User
)

type User struct {
	Username string
	Folders  []*Folder
}

func (u *User) CreateFolder(folderName string, description string) error {
	if !utils.ValidateString(folderName) {
		return fmt.Errorf("Error: %s contain invalid chars.", folderName)
	}

	if u.CheckFolder(folderName) {
		return fmt.Errorf("Error: The %s has already existed.", folderName)
	}

	folder := &Folder{
		Name:        folderName,
		CreatedAt:   time.Now().Format("yyyy-mm-dd"),
		Description: description,
	}

	u.Folders = append(u.Folders, folder)

	return nil
}

func (u *User) DeleteFolder(folderName string) error {
	if !u.CheckFolder(folderName) {
		return fmt.Errorf("Error: The %s doesnt exist.", folderName)
	}

	for i, folder := range u.Folders {
		if folder.Name == folderName {
			u.Folders = append(u.Folders[:i], u.Folders[i+1:]...)
			return nil
		}
	}

	return nil
}

func (u *User) ListFolders(sortBy string, sortOrder string) ([]*Folder, error) {
	switch sortBy {
	case "--sort--name":
		sort.Slice(u.Folders, func(i, j int) bool {
			return u.Folders[i].Name < u.Folders[j].Name
		})
	case "--sort--created":
		sort.Slice(u.Folders, func(i, j int) bool {
			return u.Folders[i].CreatedAt < u.Folders[j].CreatedAt
		})
	}

	switch sortOrder {
	case "asc":
		sort.Slice(u.Folders, func(i, j int) bool {
			return u.Folders[i].CreatedAt < u.Folders[j].CreatedAt
		})
	case "desc":
		sort.Slice(u.Folders, func(i, j int) bool {
			return u.Folders[i].CreatedAt > u.Folders[j].CreatedAt
		})
	}

	return u.Folders, nil
}

func (u *User) CheckFolder(folderName string) bool {
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

	return nil, fmt.Errorf("Error: The %s doesnt exist.", username)
}

func CheckUser(username string) bool {
	for _, user := range ListUser {
		if user.Username == username {
			return true
		}
	}

	return false
}
