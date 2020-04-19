package src

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

func readEnv(name string) string {
	return os.Getenv(name)
}

// GetPaths -- getting paths struct
func GetPaths(dir string) Paths {
	pkgPath := fmt.Sprintf("%s/package.json", dir)
	return Paths{
		DirPath:  path.Join(dir),
		PackPath: path.Join(pkgPath),
	}
}


/**
Getting extension from file name
input some.js -> returns js
*/
func getExtension(filePath string) string {
	ss := strings.Split(filePath, ".")
	s := ss[len(ss)-1]
	return s
}

/**
check if file extension should be considered or not
*/
func checkIfConsider(filePath string, supportedExtensions []string) bool {
	extension := getExtension(filePath)
	for _, v := range supportedExtensions {
		if v == extension {
			return true
		}
	}
	return false
}

// checking if the package is being used in given data or not
func checkIfDepUsed(data string, dep string) bool {
	conditions := getConditions(dep)
	for _, condition := range conditions {
		if strings.Contains(data, condition) {
			return true
		}
	}
	return false
}

// remove package from array of packages once it's being used
func removeUsedDep(deps *[]string, remove string) {
	copyArray := *deps
	var newArray []string
	for _, v := range copyArray {
		if v != remove {
			newArray = append(newArray, v)
		}
	}
	*deps = newArray
}

// get installed packages from package.json
func getInstalledDependencies(pkgData []byte) ([]string, error) {
	result := []string{}
	var pkgRaw map[string]interface{}
	if err := json.Unmarshal(pkgData, &pkgRaw); err != nil {
		return result, err
	}
	dependencies, okDep := pkgRaw["dependencies"].(map[string]interface{})
	if okDep == false {
		return result, fmt.Errorf("error  %T while type casting", pkgRaw["dependencies"])
	}
	for k := range dependencies {
		result = append(result, k)
	}
	return result, nil
}

// get file paths
func getFilePaths(paths Paths) ([]string, error) {
	var filePaths []string
	if err := getAllPaths(paths.DirPath, &filePaths, supportedExtensions); err != nil {
		return []string{}, err
	}
	return filePaths, nil
}

/**
 checking for following conditions
  1. require('package')
  2. require("package")
  3. from "package"
  4. from 'package'
*/
func getConditions(packageName string) []string {
	return []string{
		fmt.Sprintf("require('%s')", packageName),
		fmt.Sprintf("require(\"%s\")", packageName),
		fmt.Sprintf("from '%s'", packageName),
		fmt.Sprintf("from \"%s\"", packageName),
	}
}

/**
returns true if file needs to be ignored
*/
func ignoreFiles(file string) bool {
	if file == nodeModules {
		return true
	}
	return false
}