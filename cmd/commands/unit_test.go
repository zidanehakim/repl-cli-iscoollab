package commands

import (
	"fmt"
	"repl-cli-iscoollab/internal/user"
	"testing"
)

func Test_Register(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected error
	}{
		{"Valid registration", []string{"testuser"}, nil},
		{"Valid registration with space", []string{`"test user"`}, nil},
		{"Valid registration with uppercase", []string{"TestUser123"}, nil},
		{"Invalid args count (too many)", []string{"testuser", "extra"}, fmt.Errorf(user.CommandsUsage["register"])},
		{"Empty username", []string{""}, fmt.Errorf("the  contain invalid chars")},
		{"Username with spaces", []string{"test user"}, fmt.Errorf("the test user contain invalid chars")},
		{"Username with special characters", []string{"test@user"}, fmt.Errorf("the test@user contain invalid chars")},
		{"Username too long", []string{"averylongusernamethatexceedsthemaximumlength"}, fmt.Errorf("username is too long, max length allowed is 25")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Register(tt.args)
			if (err != nil) && (tt.expected != nil) {
				if err.Error() != tt.expected.Error() {
					t.Errorf("Register() error = %v, expected %v", err, tt.expected)
				}
			}
		})
	}
}

func Test_CreateFolder(t *testing.T) {
	// Register a test user first
	Register([]string{"testuser"})

	tests := []struct {
		name     string
		args     []string
		expected error
	}{
		{"Valid folder creation", []string{"testuser", "testfolder", "description"}, nil},
		{"Valid folder creation with space description", []string{"testuser", "testfolder", `"This is description"`}, nil},
		{"Valid folder creation without description", []string{"testuser", "testfolder"}, nil},
		{"Invalid args count (too few)", []string{"testuser"}, fmt.Errorf(user.CommandsUsage["create-folder"])},
		{"Invalid args count (too many)", []string{"testuser", "testfolder", "description", "extra"}, fmt.Errorf(user.CommandsUsage["create-folder"])},
		{"Empty folder name", []string{"testuser", "", "description"}, fmt.Errorf("the  contain invalid chars")},
		{"Folder name with spaces", []string{"testuser", "test folder", "description"}, fmt.Errorf("the test folder contain invalid chars")},
		{"Folder name with special characters", []string{"testuser", "test@folder", "description"}, fmt.Errorf("the test@folder contain invalid chars")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateFolder(tt.args)
			if (err != nil) && (tt.expected != nil) {
				if err.Error() != tt.expected.Error() {
					t.Errorf("CreateFolder() error = %v, expected %v", err, tt.expected)
				}
			}
		})
	}
}

func Test_ListFolders(t *testing.T) {
	// Register a test user and create folders first
	Register([]string{"testuser"})
	CreateFolder([]string{"testuser", "folder1", "description1"})
	CreateFolder([]string{"testuser", "folder2", "description2"})

	tests := []struct {
		name     string
		args     []string
		expected error
	}{
		{
			"Valid list folders",
			[]string{"testuser"},
			nil,
		},
		{
			"Valid list folders with sort by name asc",
			[]string{"testuser", "--sort-name", "asc"},
			nil,
		},
		{
			"Valid list folders with sort by name desc",
			[]string{"testuser", "--sort-name", "desc"},
			nil,
		},
		{
			"Valid list folders with sort by created asc",
			[]string{"testuser", "--sort-created", "asc"},
			nil,
		},
		{
			"Valid list folders with sort by created desc",
			[]string{"testuser", "--sort-created", "desc"},
			nil,
		},
		{
			"Invalid args count (too few)",
			[]string{},
			fmt.Errorf(user.CommandsUsage["list-folders"]),
		},
		{
			"Invalid args count (too many)",
			[]string{"testuser", "--sort-name", "asc", "extra"},
			fmt.Errorf(user.CommandsUsage["list-folders"]),
		},
		{
			"Invalid sort option",
			[]string{"testuser", "--sort-invalid", "asc"},
			fmt.Errorf(user.CommandsUsage["list-folders"]),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ListFolders(tt.args)
			if (err != nil) && (tt.expected != nil) {
				if err.Error() != tt.expected.Error() {
					t.Errorf("ListFiles() error = %v, expected %v", err, tt.expected)
				}
			}
		})
	}
}

func Test_DeleteFolder(t *testing.T) {
	// Register a test user and create a folder first
	Register([]string{"testuser"})
	CreateFolder([]string{"testuser", "testfolder", "description"})

	tests := []struct {
		name     string
		args     []string
		expected error
	}{
		{"Valid delete folder", []string{"testuser", "testfolder"}, nil},
		{"Valid delete folder with space", []string{"testuser", `"test folder"`}, nil},
		{"Invalid args count (too few)", []string{"testuser"}, fmt.Errorf(user.CommandsUsage["delete-folder"])},
		{"Invalid args count (too many)", []string{"testuser", "testfolder", "extra"}, fmt.Errorf(user.CommandsUsage["delete-folder"])},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DeleteFolder(tt.args)
			if (err != nil) && (tt.expected != nil) {
				if err.Error() != tt.expected.Error() {
					t.Errorf("DeleteFolder() error = %v, expected %v", err, tt.expected)
				}
			}
		})
	}
}

func Test_RenameFolder(t *testing.T) {
	// Register a test user and create a folder first
	Register([]string{"testuser"})
	CreateFolder([]string{"testuser", "oldfolder", "description"})

	tests := []struct {
		name     string
		args     []string
		expected error
	}{
		{"Valid rename folder", []string{"testuser", "oldfolder", "newfolder"}, nil},
		{"Invalid args count (too few)", []string{"testuser", "oldfolder"}, fmt.Errorf(user.CommandsUsage["rename-folder"])},
		{"Invalid args count (too many)", []string{"testuser", "oldfolder", "newfolder", "extra"}, fmt.Errorf(user.CommandsUsage["rename-folder"])},
		{"Empty old folder name", []string{"testuser", "", "newfolder"}, fmt.Errorf("the  doesn't exist")},
		{"Empty new folder name", []string{"testuser", "oldfolder", ""}, fmt.Errorf("the oldfolder doesn't exist")},
		{"New folder name with invalid characters", []string{"testuser", "oldfolder", "new@folder"}, fmt.Errorf("the oldfolder doesn't exist")},
		{"New folder name too long", []string{"testuser", "oldfolder", string(make([]byte, 256))}, fmt.Errorf("the oldfolder doesn't exist")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RenameFolder(tt.args)
			if (err != nil) && (tt.expected != nil) {
				if err.Error() != tt.expected.Error() {
					t.Errorf("RenameFolder() error = %v, expected %v", err, tt.expected)
				}
			}
		})
	}
}

func Test_CreateFile(t *testing.T) {
	// Register a test user and create a folder first
	Register([]string{"testuser"})
	CreateFolder([]string{"testuser", "testfolder", "description"})

	tests := []struct {
		name     string
		args     []string
		expected error
	}{
		{"Valid file creation", []string{"testuser", "testfolder", "testfile", "description"}, nil},
		{"Valid file creation without description", []string{"testuser", "testfolder", "testfile"}, nil},
		{"Invalid args count (too few)", []string{"testuser", "testfolder"}, fmt.Errorf(user.CommandsUsage["create-file"])},
		{"Invalid args count (too many)", []string{"testuser", "testfolder", "testfile", "description", "extra"}, fmt.Errorf(user.CommandsUsage["create-file"])},
		{"Empty file name", []string{"testuser", "testfolder", "", "description"}, fmt.Errorf("the  contain invalid chars")},
		{"File name with invalid characters", []string{"testuser", "testfolder", "test@file", "description"}, fmt.Errorf("the test@file contain invalid chars")},
		{"File name too long", []string{"testuser", "testfolder", string(make([]byte, 256)), "description"}, fmt.Errorf("the %s contain invalid chars", string(make([]byte, 256)))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateFile(tt.args)
			if (err != nil) && (tt.expected != nil) {
				if err.Error() != tt.expected.Error() {
					t.Errorf("CreateFile() error = %v, expected %v", err, tt.expected)
				}
			}
		})
	}
}

func Test_ListFiles(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedOutput []*user.File
		expected       error
	}{
		{"Valid list files", []string{"testuser", "testfolder"}, []*user.File{
			{Name: "file1", Description: "description1"},
			{Name: "file2", Description: "description2"},
		}, nil},
		{"Valid list files with sort by name", []string{"testuser", "testfolder", "--sort-name"}, []*user.File{
			{Name: "file1", Description: "description1"},
			{Name: "file2", Description: "description2"},
		}, nil},
		{"Valid list files with sort by created", []string{"testuser", "testfolder", "--sort-created"}, []*user.File{
			{Name: "file1", Description: "description1"},
			{Name: "file2", Description: "description2"},
		}, nil},
		{"Valid list files with sort order", []string{"testuser", "testfolder", "--sort-name", "desc"}, []*user.File{
			{Name: "file2", Description: "description2"},
			{Name: "file1", Description: "description1"},
		}, nil},
		{"Invalid args count (too few)", []string{"testuser"}, nil, fmt.Errorf(user.CommandsUsage["list-files"])},
		{"Invalid args count (too many)", []string{"testuser", "testfolder", "--sort-name", "asc", "extra"}, nil, fmt.Errorf(user.CommandsUsage["list-files"])},
		{"Invalid sort option", []string{"testuser", "testfolder", "--sort-invalid"}, nil, fmt.Errorf(user.CommandsUsage["list-files"])},
		{"Invalid sort order", []string{"testuser", "testfolder", "--sort-name", "invalid"}, nil, fmt.Errorf(user.CommandsUsage["list-files"])},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ListFiles(tt.args)
			if (err != nil) && (tt.expected != nil) {
				if err.Error() != tt.expected.Error() {
					t.Errorf("ListFiles() error = %v, expected %v", err, tt.expected)
				}
			}
		})
	}
}

func Test_DeleteFile(t *testing.T) {
	Register([]string{"testuser"})
	CreateFolder([]string{"testuser", "testfolder", "description"})
	CreateFile([]string{"testuser", "testfolder", "testfile", "description"})

	tests := []struct {
		name     string
		args     []string
		expected error
	}{
		{"Valid delete file", []string{"testuser", "testfolder", "testfile"}, nil},
		{"Valid delete file with space", []string{"testuser", "testfolder", `"test file"`}, nil},
		{"Invalid args count (too few)", []string{"testuser", "testfolder"}, fmt.Errorf(user.CommandsUsage["delete-file"])},
		{"Invalid args count (too many)", []string{"testuser", "testfolder", "testfile", "extra"}, fmt.Errorf(user.CommandsUsage["delete-file"])},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DeleteFile(tt.args)
			if (err != nil) && (tt.expected != nil) {
				if err.Error() != tt.expected.Error() {
					t.Errorf("DeleteFile() error = %v, expected %v", err, tt.expected)
				}
			}
		})
	}
}

func Test_Help(t *testing.T) {
	Help()
}
