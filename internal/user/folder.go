package user

import (
	"fmt"
	"repl-cli-iscoollab/internal/utils"
	"sort"
	"time"
)

type Folder struct {
	Name        string
	Description string
	CreatedAt   string
	Files       []*File
}

type File struct {
	Name        string
	CreatedAt   string
	Description string
}

func (f *Folder) CreateFile(fileName string, description string) error {
	if !utils.ValidateString(fileName) {
		return fmt.Errorf("Error: %s contain invalid chars.", fileName)
	}

	if f.CheckFile(fileName) {
		return fmt.Errorf("Error: The %s has already existed.", fileName)
	}

	file := &File{
		Name:        fileName,
		Description: description,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	f.Files = append(f.Files, file)

	return nil
}

func (f *Folder) DeleteFile(fileName string) error {
	for i, file := range f.Files {
		if file.Name == fileName {
			f.Files = append(f.Files[:i], f.Files[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Error: The %s doesn't exist.", fileName)
}

func (f *Folder) CheckFile(fileName string) bool {
	for _, file := range f.Files {
		if file.Name == fileName {
			return true
		}
	}
	return false
}

func (f *Folder) ListFiles(sortBy string, sortOrder string) ([]*File, error) {
	var isAsc bool
	switch sortOrder {
	case "asc":
		isAsc = true
	case "desc":
		isAsc = false
	default:
		return nil, fmt.Errorf(CommandsUsage["list-files"])
	}

	switch sortBy {
	case "--sort-name":
		sort.Slice(f.Files, func(i, j int) bool {
			if isAsc {
				return f.Files[i].Name < f.Files[j].Name
			}
			return f.Files[i].Name > f.Files[j].Name
		})
	case "--sort-created":
		sort.Slice(f.Files, func(i, j int) bool {
			if isAsc {
				return f.Files[i].CreatedAt < f.Files[j].CreatedAt
			}
			return f.Files[i].CreatedAt > f.Files[j].CreatedAt
		})
	default:
		return nil, fmt.Errorf(CommandsUsage["list-files"])
	}

	return f.Files, nil
}
