package fileStorage

import "mime/multipart"

type FileStorage struct {
	disk Disk
}

type Disk interface {
	CopyToRoot(file *multipart.FileHeader) (string, error)
	Url(path string) string
}

func NewFileSystem(disk Disk) *FileStorage {
	return &FileStorage{
		disk: disk,
	}
}

func (u *FileStorage) Store(file *multipart.FileHeader) string {
	path, err := u.disk.CopyToRoot(file)
	if err != nil {
		return ""
	}

	return path
}

func (u FileStorage) Url(filePath string) string {
	return u.disk.Url(filePath)
}