package file

import (
	"os"
)

type File struct {
	path string
	file *os.File
}

func CreateFile(path string) (*File, error) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.FileMode(0644))
	if err != nil {
		return nil, err
	}

	return &File{
		path: path,
		file: file,
	}, nil
}

func (f *File) CloseFile() {
	f.file.Close()
}
