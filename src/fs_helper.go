package src

import (
	"bufio"
	"io/ioutil"
	"os"
	"path"
)

type scanner struct {
	file *os.File
	reader *bufio.Scanner
}

func readData(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func getAllPaths(dirPath string, paths *[]string, supportedExtensions []string) error {
	fileInfo, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, file := range fileInfo {
		if file.IsDir() {
			if ignoreFiles(file.Name()) {
				continue
			}
			nextPath := path.Join(dirPath, file.Name())
			if err := getAllPaths(nextPath, paths, supportedExtensions); err != nil {
				return err
			}
		} else {
			extension := file.Name()
			if checkIfConsider(extension, supportedExtensions) {
				nextPath := path.Join(dirPath, extension)
				*paths = append(*paths, nextPath)
			}
		}
	}
	return nil
}
