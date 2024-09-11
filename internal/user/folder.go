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
	Files       map[string]*File
}

type File struct {
	Name        string
	CreatedAt   string
	Description string
}

func (f *Folder) CreateFile(fileName string, description string) error {
	if !utils.ValidateString(fileName) {
		return fmt.Errorf("the %s contain invalid chars", fileName)
	}

	if _, exists := f.Files[fileName]; exists {
		return fmt.Errorf("the %s has already existed", fileName)
	}

	if len(fileName) > MaxFileNameLength {
		return fmt.Errorf("filename is too long, max length allowed is %d", MaxFileNameLength)
	}

	file := &File{
		Name:        fileName,
		Description: description,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	f.Files[fileName] = file

	return nil
}

func (f *Folder) DeleteFile(fileName string) error {
	if _, exists := f.Files[fileName]; exists {
		delete(f.Files, fileName)
		return nil
	}

	return fmt.Errorf("the %s doesn't exist", fileName)
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

	files := make([]*File, 0, len(f.Files))
	for _, file := range f.Files {
		files = append(files, file)
	}

	switch sortBy {
	case "--sort-name":
		sort.Slice(files, func(i, j int) bool {
			if isAsc {
				return files[i].Name < files[j].Name
			}
			return files[i].Name > files[j].Name
		})
	case "--sort-created":
		sort.Slice(files, func(i, j int) bool {
			if isAsc {
				return files[i].CreatedAt < files[j].CreatedAt
			}
			return files[i].CreatedAt > files[j].CreatedAt
		})
	default:
		return nil, fmt.Errorf(CommandsUsage["list-files"])
	}

	return files, nil
}
