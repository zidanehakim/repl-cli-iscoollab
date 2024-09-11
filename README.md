# Virtual File System - REPL CLI

![Go Version](https://img.shields.io/badge/go-1.23.1-blue.svg)

## 📖 Overview

This project implements a **Virtual File System** using a REPL (Read-Eval-Print Loop) interface in Go (version 1.23.1). The VFS allows users to manage users, folders, and files in memory, without persistent storage. It provides a command-line interface to interact with the system, performing actions like user registration, folder creation, and file management.

This project is a task assignment for the IsCoolLab Backend Engineer Intern position.

## 🌟 Features

### 👤 User Management
- Register unique, case-insensitive usernames
- Handle multiple folders and files for each user

### 📁 Folder Management
- Create, delete, and rename folders
- Case-insensitive folder names (unique within a user's scope)
- Optional folder description field

### 📄 File Management
- Create, delete, and list files in specific folders
- Case-insensitive file names (unique within each folder)
- Optional file description field

## 🚀 Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/zidanehakim/repl-cli-iscoollab.git
   cd repl-cli-iscoollab
   ```

2. **Build the project:**

   ```bash
   go build -o [appname]
   ```

3. **Run the executable:**

   ```bash
   ./[appname]
   ```

   Note: You can replace "[appname]" with any name you prefer for your application.

## 🖥️ Usage

The application runs as a REPL interface. Upon starting, you can input various commands to interact with the virtual file system.

### 🛠️ Commands

#### User Registration

- **Command**: `register [username]`
- **Example**: `register john_doe`
- **Responses**:
  - **Success**: `Add [username] successfully`
  - **Error**: `the [username] has already existed` or `the [username] contains invalid chars`

#### Folder Management

- **Create Folder**:
  - **Command**: `create-folder [username] [foldername] [description]?`
  - **Example**: `create-folder john_doe my_folder "This is my folder"`
  - **Success**: `Create [foldername] successfully`
  - **Error**: `the [username] doesn't exist` or `the [foldername] contains invalid chars`

- **Delete Folder**:
  - **Command**: `delete-folder [username] [foldername]`
  - **Example**: `delete-folder john_doe my_folder`
  - **Success**: `Delete [foldername] successfully`
  - **Error**: `the [foldername] doesn't exist`

- **List Folders**:
  - **Command**: `list-folders [username] [--sort-name|--sort-created] [asc|desc]`
  - **Example**: `list-folders john_doe --sort-name asc`
  - Displays a list of folders.
  - **Error**: `the [username] doesn't exist`

- **Rename Folder**:
  - **Command**: `rename-folder [username] [foldername] [new-folder-name]`
  - **Example**: `rename-folder john_doe my_folder new_folder`
  - **Success**: `Rename [foldername] to [new-folder-name] successfully`

#### File Management

- **Create File**:
  - **Command**: `create-file [username] [foldername] [filename] [description]?`
  - **Example**: `create-file john_doe my_folder my_file "This is my file"`
  - **Success**: `Create [filename] in [username]/[foldername] successfully`
  - **Error**: `the [username] doesn't exist` or `The [filename] contains invalid chars`

- **Delete File**:
  - **Command**: `delete-file [username] [foldername] [filename]`
  - **Example**: `delete-file john_doe my_folder my_file`
  - **Success**: `Delete [filename] successfully`

- **List Files**:
  - **Command**: `list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]`
  - **Example**: `list-files john_doe my_folder --sort-name asc`
  - Lists files in the specified folder.

  For a full list of available commands, type `help` at the prompt.

### ✅ Input Validation

- Usernames, folder names, and file names must not contain invalid characters (e.g., `@`)
- Commands follow strict syntax; invalid commands or incorrect flags will result in an error message
- Accept words with whitespace input by putting double quote `"` or `'` in between (e.g `"New Folder"`)
- Any extra whitespace (more than one) will be simplified as one whitespace

### 📝 Example Usage

```bash
# Register two users
register user1
# Output: Add user1 successfully
register user2
# Output: Add user2 successfully

# Create folders for user1
create-folder user1 folder1
# Output: Create folder1 successfully

# Create a folder with description
create-folder user1 folder2 "My second folder"
# Output: Create folder2 successfully

# List folders for user1
list-folders user1 --sort-name asc
# Output:
# folder1 2023-01-01 15:00:00 user1
# folder2 "My second folder" 2023-01-01 15:00:10 user1

# Create a file for user1
create-file user1 folder1 file1 "My first file"
# Output: Create file1 in user1/folder1 successfully

# List files in a folder
list-files user1 folder1 --sort-name asc
# Output:
# file1 "My first file" 2023-01-01 15:00:20 folder1 user1

# Invalid command
invalid-command
# Output: Error: Unrecognized command
```

## 🏗️ Project Architecture

The project follows Go's file structure conventions, organizing the code into packages for clean separation of concerns:

```
repl-cli-iscoollab/
├── cmd/
│   └── commands/
│       └── commands.go
|       └── unit_test.go
├── internal/
│   ├── user/
│   │   └── user.go
│   |   └── folder.go
│   └── utils/
│       └── utils.go
├── main.go
├── integration_test.go
└── go.mod
```

- **`main.go`**: Entry point for the application
- **`cmd/`**: Contains CLI-related code
- **`internal/`**: Houses core logic and data management
- **`utils/`**: Helper functions
- **`user/`**: Contains global variable for program's memory state

### Data Management
I chose to use `map` rather than arrays. This decision is based on the need for efficient lookups and quick access to user, folder, and file data. `map` provides O(1) average time complexity for lookups, which is crucial for performance in this project. Although arrays could be used in some scenarios, the dynamic nature of the data (frequent insertions and deletions) made `map` a more suitable choice.

## 📄 License

This project is not licensed. All rights reserved. 