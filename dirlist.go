package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//cd C:\Users\rober\go
//go run src\github.com\Rob8150\dirlist\directory.go

func dirlist(searchDir string, match string) ([]string, error) {

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, match) {
			fileList = append(fileList, path)
		}

		return err
	})

	if e != nil {
		panic(e)
	}

	for _, file := range fileList {
		fmt.Println(file)
	}

	return fileList, nil

}

func main() {
	search := "github.com/Rob8150"
	match := ".go"
	dirlist(search, match)
}
