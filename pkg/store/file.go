package store

import (
	"encoding/json"
	"errors"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string

type FileStore struct {
	FileName string
}

const (
	FileType Type = "filestorage"
)

func (fs *FileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	f, err := os.OpenFile(fs.FileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(fileData)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		file = []byte("[]")
	}
	return json.Unmarshal(file, &data)
}

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{fileName}
	}
	return nil
}
