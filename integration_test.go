package main

import (
	"repl-cli-iscoollab/cmd/commands"
	"strings"
	"testing"
	"time"
)

func Test_Integration(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Register user1", "register user1", "Add user1 successfully"},
		{"Register user2", "register user2", "Add user2 successfully"},
		{"Create folder1 for user1", "create-folder user1 folder1", "Create folder1 successfully"},
		{"Create folder1 for user2", "create-folder user2 folder1", "Create folder1 successfully"},
		{"Attempt to create existing folder", "create-folder user1 folder1", "Error: the folder1 has already existed"},
		{"Create folder2 with description for user1", "create-folder user1 folder2 this-is-folder-2", "Create folder2 successfully"},
		{"List folders for user1 sorted by name", "list-folders user1 --sort-name asc", "folder1 " + time.Now().Format("2006-01-02 15:04:05") + " user1\nfolder2 this-is-folder-2 " + time.Now().Format("2006-01-02 15:04:05") + " user1"},
		{"List folders for user2", "list-folders user2", "folder1 " + time.Now().Format("2006-01-02 15:04:05") + " user2"},
		{"Create file1 for user1 in folder1", "create-file user1 folder1 file1 this-is-file1", "Create file1 in user1/folder1 successfully"},
		{"Create config file for user1 in folder1", "create-file user1 folder1 config a-config-file", "Create config in user1/folder1 successfully"},
		{"Attempt to create existing file", "create-file user1 folder1 config a-config-file", "Error: the config has already existed"},
		{"Attempt to create file for unregistered user", "create-file user-abc folder-abc config a-config-file", "Error: the user-abc doesn't exist"},
		{"Attempt unsupported command", "list data", "Error: Unrecognized command"},
		{"Attempt to list files with incorrect flags", "list-files user1 folder1 --sort a", "Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]"},
		{"List files sorted by name desc", "list-files user1 folder1 --sort-name desc", "file1 this-is-file1 " + time.Now().Format("2006-01-02 15:04:05") + " user1\nconfig a-config-file " + time.Now().Format("2006-01-02 15:04:05") + " user1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output string
			var err error

			switch {
			case strings.HasPrefix(tt.input, "register"):
				output, err = commands.Register(strings.Fields(tt.input)[1:])
			case strings.HasPrefix(tt.input, "create-folder"):
				output, err = commands.CreateFolder(strings.Fields(tt.input)[1:])
			case strings.HasPrefix(tt.input, "list-folders"):
				output, err = commands.ListFolders(strings.Fields(tt.input)[1:])
			case strings.HasPrefix(tt.input, "create-file"):
				output, err = commands.CreateFile(strings.Fields(tt.input)[1:])
			case strings.HasPrefix(tt.input, "list-files"):
				output, err = commands.ListFiles(strings.Fields(tt.input)[1:])
			case strings.HasPrefix(tt.input, "delete-folder"):
				output, err = commands.DeleteFolder(strings.Fields(tt.input)[1:])
			case strings.HasPrefix(tt.input, "delete-file"):
				output, err = commands.DeleteFile(strings.Fields(tt.input)[1:])
			case strings.HasPrefix(tt.input, "rename-folder"):
				output, err = commands.RenameFolder(strings.Fields(tt.input)[1:])
			case strings.HasPrefix(tt.input, "help"):
				commands.Help()
			default:
				output = "Error: Unrecognized command"
			}

			if err != nil {
				if strings.Contains(err.Error(), "Usage: ") {
					output = err.Error()
				} else {
					output = "Error: " + err.Error()
				}
			}

			if !strings.Contains(output, tt.expected) {
				t.Errorf("Test %s failed. Expected: %s, Got: %s", tt.name, tt.expected, output)
			}
		})
	}
}
