package user

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
