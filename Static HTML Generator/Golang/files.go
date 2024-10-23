package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// --------------------------------------------------------------------------------------------------------------------

// Copy data

// --------------------------------------------------------------------------------------------------------------------

// Copy files preserving folder structure.

func copyFiles(rootDir string) error {
	filePaths, err := getFileNames(rootDir)
	if err != nil {
		return err
	}

	var newPaths []string
	for _, path := range filePaths {
		newPath := strings.Replace(path, "static", "public", 1)
		newPaths = append(newPaths, newPath)
	}

	for i, path := range filePaths {
		err = copyCurrentFile(path, newPaths[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// Gets all the sub-paths from a starting path.

func getFileNames(rootDir string) ([]string, error) {
	currentDirectory, err := os.Getwd()
	if err != nil {
		return []string{}, fmt.Errorf("getFileNames: %v", err)
	}

	originPath := filepath.Join(currentDirectory, rootDir)
	_, err = os.Stat(originPath)
	if os.IsNotExist(err) {
		return []string{}, fmt.Errorf("getFileNames: %v", err)
	}

	var filePaths []string

	err = filepath.Walk(originPath,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			filePaths = append(filePaths, path)
			return nil
		})

	if err != nil {
		return []string{}, err
	}
	return filePaths, nil
}

// --------------------------------------------------------------------------------------------------------------------

// Copies a file or directory.

func copyCurrentFile(sourcePath, destinationPath string) error {
	file, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("copyCurrentFile: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("copyCurrentFile: %v", err)
	}

	permissionLevel := 0777
	if fileInfo.IsDir() {
		os.Mkdir(destinationPath, fs.FileMode(permissionLevel))
	} else {
		newFile, err := os.Create(destinationPath)
		if err != nil {
			return fmt.Errorf("copyCurrentFile: %v", err)
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err != nil {
			return fmt.Errorf("copyCurrentFile: %v", err)
		}
	}
	return nil
}

// --------------------------------------------------------------------------------------------------------------------
