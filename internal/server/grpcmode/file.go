package grpcmode

import (
	"log"
	"os"
	"path/filepath"
)

// File хранит параметры файла
type File struct {
	FilePath   string
	OutputFile *os.File
}

// NewFile создает новый файл
func NewFile() *File {
	return &File{}
}

// SetFile создает файл
func (f *File) SetFile(fileName, path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	f.FilePath = filepath.Join(path, fileName)
	file, err := os.Create(f.FilePath)
	if err != nil {
		return err
	}
	f.OutputFile = file
	return nil
}

// Write хаписывает данные в файл
func (f *File) Write(chunk []byte) error {
	if f.OutputFile == nil {
		return nil
	}
	_, err := f.OutputFile.Write(chunk)
	return err
}

// Close закрывает файл
func (f *File) Close() error {
	return f.OutputFile.Close()
}
