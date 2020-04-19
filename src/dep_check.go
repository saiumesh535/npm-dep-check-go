package src

import (
	"fmt"
)

// DepCheck -- checking dep
func DepCheck() error {
	// reading env
	dir := readEnv(DIR)
	if dir == "" {
		return fmt.Errorf("%s cannot be empty", DIR)
	}

	// getting paths struct
	paths := GetPaths(dir)

	// reading data from package json
	pkgData, err := readData(paths.PackPath)
	if err != nil {
		return err
	}

	// get all installed dependencies
	packages, err := getInstalledDependencies(pkgData)
	if err != nil {
		return err
	}

	filePaths, err := getFilePaths(paths)
	if err != nil {
		return err
	}

	// assume all the packages or not uses
	unUsedPackages := append([]string{}, packages...)

	for _, filePath := range filePaths {
		data, err := readData(filePath)
		if err != nil {
			return err
		}
		dataInString := string(data)
		for _, unUsedPackage := range unUsedPackages {
			if checkIfDepUsed(dataInString, unUsedPackage) {
				removeUsedDep(&unUsedPackages, unUsedPackage)
			}
		}
	}
	fmt.Println(unUsedPackages)
	return nil
}
