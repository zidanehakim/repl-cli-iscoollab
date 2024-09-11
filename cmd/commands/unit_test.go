package commands

import (
	"fmt"
	"repl-cli-iscoollab/internal/user"
	"testing"
	"time"
)

// Test_Register tests the Register function with various input scenarios.
// Testing strategy:
// 1. Test valid registrations (normal, with space, with uppercase)
// 2. Test invalid registrations (too many args, empty username, invalid characters, too long)
// 3. Test registering an existing user
func Test_Register(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedOutput string
		expectedError  error
	}{
		{"Valid registration", []string{"testuser"}, "Add testuser successfully\n", nil},
		{"Valid registration with space", []string{`"test user"`}, "Add \"test user\" successfully\n", nil},
		{"Valid registration with uppercase", []string{"TestUser123"}, "Add testuser123 successfully\n", nil},
		{"Invalid args count (too many)", []string{"testuser", "extra"}, "", fmt.Errorf(user.CommandsUsage["register"])},
		{"Empty username", []string{""}, "", fmt.Errorf("the  contain invalid chars")},
		{"Username with spaces", []string{"test user"}, "", fmt.Errorf("the test user contain invalid chars")},
		{"Username with special characters", []string{"test@user"}, "", fmt.Errorf("the test@user contain invalid chars")},
		{"Username too long", []string{"averylongusernamethatexceedsthemaximumlength"}, "", fmt.Errorf("username is too long, max length allowed is 25")},
		{"Nonexistent user", []string{"nonexistentuser"}, "", fmt.Errorf("the nonexistentuser has already existed")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Nonexistent user" {
				Register([]string{"nonexistentuser"})
			}
			output, err := Register(tt.args)
			if (err != nil) != (tt.expectedError != nil) {
				t.Errorf("Register() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("Register() error = %v, expectedError %v", err, tt.expectedError)
			}
			if output != tt.expectedOutput {
				t.Errorf("Register() output = %v, expectedOutput %v", output, tt.expectedOutput)
			}
		})
	}
}

// Test_CreateFolder tests the CreateFolder function with various input scenarios.
// Testing strategy:
// 1. Test valid folder creation (normal, with space in description)
// 2. Test invalid folder creation (too few/many args, empty folder name, invalid characters)
// 3. Test creating an existing folder
func Test_CreateFolder(t *testing.T) {
	// Register a test user first
	Register([]string{"testuser"})

	tests := []struct {
		name           string
		args           []string
		expectedOutput string
		expectedError  error
	}{
		{"Valid folder creation", []string{"testuser", "testfolder", "description"}, "Create testfolder successfully\n", nil},
		{"Valid folder creation with space description", []string{"testuser", "\"test folder\"", `"This is description"`}, "Create \"test folder\" successfully\n", nil},
		{"Invalid args count (too few)", []string{"testuser"}, "", fmt.Errorf(user.CommandsUsage["create-folder"])},
		{"Invalid args count (too many)", []string{"testuser", "testfolder", "description", "extra"}, "", fmt.Errorf(user.CommandsUsage["create-folder"])},
		{"Empty folder name", []string{"testuser", "", "description"}, "", fmt.Errorf("the  contain invalid chars")},
		{"Folder name with spaces", []string{"testuser", "test folder", "description"}, "", fmt.Errorf("the test folder contain invalid chars")},
		{"Folder name with special characters", []string{"testuser", "test@folder", "description"}, "", fmt.Errorf("the test@folder contain invalid chars")},
		{"Nonexistent folder", []string{"testuser", "nonexistentfolder", "description"}, "", fmt.Errorf("the nonexistentfolder has already existed")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Nonexistent folder" {
				CreateFolder([]string{"testuser", "nonexistentfolder", "description"})
			}
			output, err := CreateFolder(tt.args)
			if (err != nil) != (tt.expectedError != nil) {
				t.Errorf("CreateFolder() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("CreateFolder() error = %v, expectedError %v", err, tt.expectedError)
			}
			if output != tt.expectedOutput {
				t.Errorf("CreateFolder() output = %v, expectedOutput %v", output, tt.expectedOutput)
			}
		})
	}
}

// Test_ListFolders tests the ListFolders function with various input scenarios.
// Testing strategy:
// 1. Test listing folders without sorting
// 2. Test listing folders with different sorting options (name/created, asc/desc)
// 3. Test invalid cases (too few/many args, invalid sort option)
// 4. Test listing folders for a nonexistent user
func Test_ListFolders(t *testing.T) {
	// Register a test user and create folders first
	Register([]string{"testuser"})
	CreateFolder([]string{"testuser", "folder1", "description1"})
	CreateFolder([]string{"testuser", "folder2", "description2"})
	DeleteFolder([]string{"testuser", "testfolder"})
	DeleteFolder([]string{"testuser", "\"test folder\""})
	DeleteFolder([]string{"testuser", "nonexistentfolder"})

	tests := []struct {
		name           string
		args           []string
		expectedOutput string
		expectedError  error
	}{
		{
			"Valid list folders",
			[]string{"testuser"},
			fmt.Sprintf("folder1 description1 %s testuser\nfolder2 description2 %s testuser\n", time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05")),
			nil,
		},
		{
			"Valid list folders with sort by name asc",
			[]string{"testuser", "--sort-name", "asc"},
			fmt.Sprintf("folder1 description1 %s testuser\nfolder2 description2 %s testuser\n", time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05")),
			nil,
		},
		{
			"Valid list folders with sort by name desc",
			[]string{"testuser", "--sort-name", "desc"},
			fmt.Sprintf("folder2 description2 %s testuser\nfolder1 description1 %s testuser\n", time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05")),
			nil,
		},
		{
			"Valid list folders with sort by created asc",
			[]string{"testuser", "--sort-created", "asc"},
			fmt.Sprintf("folder1 description1 %s testuser\nfolder2 description2 %s testuser\n", time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05")),
			nil,
		},
		{
			"Valid list folders with sort by created desc",
			[]string{"testuser", "--sort-created", "desc"},
			fmt.Sprintf("folder1 description1 %s testuser\nfolder2 description2 %s testuser\n", time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05")),
			nil,
		},
		{
			"Invalid args count (too few)",
			[]string{},
			"",
			fmt.Errorf(user.CommandsUsage["list-folders"]),
		},
		{
			"Invalid args count (too many)",
			[]string{"testuser", "--sort-name", "asc", "extra"},
			"",
			fmt.Errorf(user.CommandsUsage["list-folders"]),
		},
		{
			"Invalid sort option",
			[]string{"testuser", "--sort-invalid", "asc"},
			"",
			fmt.Errorf(user.CommandsUsage["list-folders"]),
		},
		{
			"Nonexistent user",
			[]string{"nonexistentusers"},
			"",
			fmt.Errorf("the nonexistentusers doesn't exist"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := ListFolders(tt.args)
			if (err != nil) != (tt.expectedError != nil) {
				t.Errorf("ListFolders() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("ListFolders() error = %v, expectedError %v", err, tt.expectedError)
			}
			if output != tt.expectedOutput {
				t.Errorf("ListFolders() output = %v, expectedOutput %v", output, tt.expectedOutput)
			}
		})
	}
}

// Test_DeleteFolder tests the DeleteFolder function with various input scenarios.
// Testing strategy:
// 1. Test valid folder deletion (normal, with space in folder name)
// 2. Test invalid folder deletion (too few/many args)
// 3. Test deleting a nonexistent folder
func Test_DeleteFolder(t *testing.T) {
	// Register a test user and create a folder first
	Register([]string{"testuser"})
	CreateFolder([]string{"testuser", "testfolder", "description"})
	CreateFolder([]string{"testuser", "\"test folder\"", "description"})

	tests := []struct {
		name           string
		args           []string
		expectedOutput string
		expectedError  error
	}{
		{"Valid delete folder", []string{"testuser", "testfolder"}, "Delete testfolder successfully\n", nil},
		{"Valid delete folder with space", []string{"testuser", `"test folder"`}, "Delete \"test folder\" successfully\n", nil},
		{"Invalid args count (too few)", []string{"testuser"}, "", fmt.Errorf(user.CommandsUsage["delete-folder"])},
		{"Invalid args count (too many)", []string{"testuser", "testfolder", "extra"}, "", fmt.Errorf(user.CommandsUsage["delete-folder"])},
		{"Nonexistent folder", []string{"testuser", "nonexistentfolder"}, "", fmt.Errorf("the nonexistentfolder doesn't exist")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Nonexistent folder" {
				DeleteFolder([]string{"testuser", "nonexistentfolder"})
			}
			output, err := DeleteFolder(tt.args)
			if (err != nil) != (tt.expectedError != nil) {
				t.Errorf("DeleteFolder() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("DeleteFolder() error = %v, expectedError %v", err, tt.expectedError)
			}
			if output != tt.expectedOutput {
				t.Errorf("DeleteFolder() output = %v, expectedOutput %v", output, tt.expectedOutput)
			}
		})
	}
}

// Test_RenameFolder tests the RenameFolder function with various input scenarios.
// Testing strategy:
// 1. Test valid folder renaming
// 2. Test invalid folder renaming (too few/many args, empty old/new folder names, invalid characters, too long name)
// 3. Test renaming a nonexistent folder
func Test_RenameFolder(t *testing.T) {
	// Register a test user and create a folder first
	Register([]string{"testuser"})
	CreateFolder([]string{"testuser", "oldfolder", "description"})

	tests := []struct {
		name           string
		args           []string
		expectedOutput string
		expectedError  error
	}{
		{"Valid rename folder", []string{"testuser", "oldfolder", "newfolder"}, "Rename oldfolder to newfolder successfully\n", nil},
		{"Invalid args count (too few)", []string{"testuser", "oldfolder"}, "", fmt.Errorf(user.CommandsUsage["rename-folder"])},
		{"Invalid args count (too many)", []string{"testuser", "oldfolder", "newfolder", "extra"}, "", fmt.Errorf(user.CommandsUsage["rename-folder"])},
		{"Empty old folder name", []string{"testuser", "", "newfolder"}, "", fmt.Errorf("the  doesn't exist")},
		{"Empty new folder name", []string{"testuser", "oldfolder", ""}, "", fmt.Errorf("the oldfolder doesn't exist")},
		{"New folder name with invalid characters", []string{"testuser", "oldfolder", "new@folder"}, "", fmt.Errorf("the oldfolder doesn't exist")},
		{"New folder name too long", []string{"testuser", "oldfolder", string(make([]byte, 256))}, "", fmt.Errorf("the oldfolder doesn't exist")},
		{"Nonexistent folder", []string{"testuser", "nonexistentfolder", "newfolder"}, "", fmt.Errorf("the nonexistentfolder doesn't exist")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Nonexistent folder" {
				DeleteFolder([]string{"testuser", "nonexistentfolder"})
			}
			output, err := RenameFolder(tt.args)
			if (err != nil) != (tt.expectedError != nil) {
				t.Errorf("RenameFolder() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("RenameFolder() error = %v, expectedError %v", err, tt.expectedError)
			}
			if output != tt.expectedOutput {
				t.Errorf("RenameFolder() output = %v, expectedOutput %v", output, tt.expectedOutput)
			}
		})
	}
}

// Test_CreateFile tests the CreateFile function with various input scenarios.
// Testing strategy:
// 1. Test valid file creation
// 2. Test invalid file creation (too few/many args, empty file name, invalid characters, too long name)
// 3. Test creating a file in a nonexistent folder
func Test_CreateFile(t *testing.T) {
	// Register a test user and create a folder first
	Register([]string{"testuser"})
	CreateFolder([]string{"testuser", "testfolder", "description"})

	tests := []struct {
		name           string
		args           []string
		expectedOutput string
		expectedError  error
	}{
		{"Valid file creation", []string{"testuser", "testfolder", "testfile", "description"}, "Create testfile in testuser/testfolder successfully\n", nil},
		{"Invalid args count (too few)", []string{"testuser", "testfolder"}, "", fmt.Errorf(user.CommandsUsage["create-file"])},
		{"Invalid args count (too many)", []string{"testuser", "testfolder", "testfile", "description", "extra"}, "", fmt.Errorf(user.CommandsUsage["create-file"])},
		{"Empty file name", []string{"testuser", "testfolder", "", "description"}, "", fmt.Errorf("the  contain invalid chars")},
		{"File name with invalid characters", []string{"testuser", "testfolder", "test@file", "description"}, "", fmt.Errorf("the test@file contain invalid chars")},
		{"File name too long", []string{"testuser", "testfolder", string(make([]byte, 256)), "description"}, "", fmt.Errorf("the %s contain invalid chars", string(make([]byte, 256)))},
		{"Nonexistent folder", []string{"testuser", "nonexistentfolder", "testfile", "description"}, "", fmt.Errorf("the nonexistentfolder doesn't exist")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Nonexistent folder" {
				DeleteFolder([]string{"testuser", "nonexistentfolder"})
			}
			output, err := CreateFile(tt.args)
			if (err != nil) != (tt.expectedError != nil) {
				t.Errorf("CreateFile() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("CreateFile() error = %v, expectedError %v", err, tt.expectedError)
			}
			if output != tt.expectedOutput {
				t.Errorf("CreateFile() output = %v, expectedOutput %v", output, tt.expectedOutput)
			}
		})
	}
}

// Test_ListFiles tests the ListFiles function with various input scenarios.
// Testing strategy:
// 1. Test valid file listing (normal, with sorting options)
// 2. Test invalid file listing (too few/many args, invalid sort option)
// 3. Test listing files in a nonexistent folder
func Test_ListFiles(t *testing.T) {
	// Register a test user and create a folder with files first
	Register([]string{"testuser"})
	CreateFolder([]string{"testuser", "testfolder", "description"})
	CreateFile([]string{"testuser", "testfolder", "file1", "description1"})
	CreateFile([]string{"testuser", "testfolder", "file2", "description2"})
	DeleteFile([]string{"testuser", "testfolder", "testfile"})

	now := time.Now().Format("2006-01-02 15:04:05")
	tests := []struct {
		name           string
		args           []string
		expectedOutput string
		expectedError  error
	}{
		{"Valid list files", []string{"testuser", "testfolder"}, fmt.Sprintf("file1 description1 %s testuser\nfile2 description2 %s testuser\n", now, now), nil},
		{"Valid list files with sort by name", []string{"testuser", "testfolder", "--sort-name", "asc"}, fmt.Sprintf("file1 description1 %s testuser\nfile2 description2 %s testuser\n", now, now), nil},
		{"Valid list files with sort by created", []string{"testuser", "testfolder", "--sort-created", "asc"}, fmt.Sprintf("file1 description1 %s testuser\nfile2 description2 %s testuser\n", now, now), nil},
		{"Valid list files with sort order desc", []string{"testuser", "testfolder", "--sort-name", "desc"}, fmt.Sprintf("file2 description2 %s testuser\nfile1 description1 %s testuser\n", now, now), nil},
		{"Invalid args count (too few)", []string{"testuser"}, "", fmt.Errorf(user.CommandsUsage["list-files"])},
		{"Invalid args count (too many)", []string{"testuser", "testfolder", "--sort-name", "asc", "extra"}, "", fmt.Errorf(user.CommandsUsage["list-files"])},
		{"Invalid sort option", []string{"testuser", "testfolder", "--sort-invalid", "asc"}, "", fmt.Errorf(user.CommandsUsage["list-files"])},
		{"Invalid sort order", []string{"testuser", "testfolder", "--sort-name", "invalid"}, "", fmt.Errorf(user.CommandsUsage["list-files"])},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := ListFiles(tt.args)
			if (err != nil) != (tt.expectedError != nil) {
				t.Errorf("ListFiles() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("ListFiles() error = %v, expectedError %v", err, tt.expectedError)
			}
			if output != tt.expectedOutput {
				t.Errorf("ListFiles() output = %v, expectedOutput %v", output, tt.expectedOutput)
			}
		})
	}
}

// Test_DeleteFile tests the DeleteFile function with various input scenarios.
// Testing strategy:
// 1. Test valid file deletion (normal, with space in file name)
// 2. Test invalid file deletion (too few/many args)
// 3. Test deleting a nonexistent file
func Test_DeleteFile(t *testing.T) {
	Register([]string{"testuser"})
	CreateFolder([]string{"testuser", "testfolder", "description"})
	CreateFile([]string{"testuser", "testfolder", "testfile", "description"})
	CreateFile([]string{"testuser", "testfolder", "\"test file\"", "description"})

	tests := []struct {
		name           string
		args           []string
		expectedOutput string
		expectedError  error
	}{
		{"Valid delete file", []string{"testuser", "testfolder", "testfile"}, "Deleted file testfile from testuser/testfolder successfully\n", nil},
		{"Valid delete file with space", []string{"testuser", "testfolder", `"test file"`}, "Deleted file \"test file\" from testuser/testfolder successfully\n", nil},
		{"Invalid args count (too few)", []string{"testuser", "testfolder"}, "", fmt.Errorf(user.CommandsUsage["delete-file"])},
		{"Invalid args count (too many)", []string{"testuser", "testfolder", "testfile", "extra"}, "", fmt.Errorf(user.CommandsUsage["delete-file"])},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := DeleteFile(tt.args)
			if (err != nil) != (tt.expectedError != nil) {
				t.Errorf("DeleteFile() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("DeleteFile() error = %v, expectedError %v", err, tt.expectedError)
			}
			if output != tt.expectedOutput {
				t.Errorf("DeleteFile() output = %v, expectedOutput %v", output, tt.expectedOutput)
			}
		})
	}
}

// Test_Help tests the Help function.
// Testing strategy:
// 1. Ensure the Help function runs without errors
func Test_Help(t *testing.T) {
	Help()
}
