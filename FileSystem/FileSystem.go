package FileSystem

import (
	"fmt"
	"io"
)

type FileType int32

const (
	_ UNKNOWNFILETYPE = iota
	CSV
	PARQUET
	ORC
)

type FileLocation struct {
	Location string
	Type     FileType
}

func NewFileLocation(loc string, ft FileType) *FileLocation {
	return &FileLocation{
		Location: loc,
		Type:     ft,
	}
}

type VirtualFile interface {
	io.ReaderAt
	io.ReadCloser
	io.Seeker
	Size() int64
}

var (
	fileSystems = []VirtualFileSystem{
		&LocalFileSystem{},
		&HdfsFileSystem{},
		&S3FileSystem{},
	}
)

type VirtualFileSystem interface {
	Accept(*FileLocation) bool
	Open(*FileLocation) (VirtualFile, error)
	List(*FileLocation) ([]*FileLocation, error)
	IsDir(*FileLocation) bool
}

func Open(filepath string) (VirtualFile, error) {
	fileLocation := &FileLocation{filepath}
	for _, fs := range fileSystems {
		if fs.Accept(fileLocation) {
			return fs.Open(fileLocation)
		}
	}
	return nil, fmt.Errorf("Unknown file %s", filepath)
}

func List(filepath string) ([]*FileLocation, error) {
	fileLocation := &FileLocation{filepath}
	for _, fs := range fileSystems {
		if fs.Accept(fileLocation) {
			return fs.List(fileLocation)
		}
	}
	return nil, fmt.Errorf("Unknown file %s", filepath)
}

func IsDir(filepath string) bool {
	fileLocation := &FileLocation{filepath}
	for _, fs := range fileSystems {
		if fs.Accept(fileLocation) {
			return fs.IsDir(fileLocation)
		}
	}
	return false
}
