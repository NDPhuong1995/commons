package filesystem

import (
	"fmt"
	"os"
)

func HasExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsExist(err)
}

func HasPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

func CreateFolderNotExist(src string) error {
	if !HasExist(src) {
		return os.MkdirAll(src, os.ModePerm)
	}
	return nil
}

func MustOpenFile(fileName, dir string) (*os.File, error) {
	perm := HasPermission(dir)
	if perm == true {
		return nil, fmt.Errorf("dir: %s permission denied ", dir)
	}
	err := CreateFolderNotExist(dir)
	if err != nil {
		return nil, fmt.Errorf("create folder dir: %s error: %v", dir, err)
	}
	f, err := os.OpenFile(dir + fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("get file path: %s error %v", dir + fileName, err)
	}
	return f, nil
}
