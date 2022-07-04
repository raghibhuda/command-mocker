package main

import (
	"fmt"
	"os"
)

func deleteASingleFile(filepath string) error {
	err := os.Remove(filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return nil
}

func deleteDirectory(dirPath string) error {
	err := os.RemoveAll(dirPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return nil
}
func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

func createSingleDir(dirName string) error {
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return nil
}

func createNestedDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return nil
}

func runPWD() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return dir
}

func runRM(args []string, recursive bool) error {
	if !recursive {
		for _, file := range args {
			isDir, _ := isDirectory(file)
			if isDir {
				fmt.Println("Skipped directory can not be deleted")
				continue
			}
			deleteASingleFile(file)
		}
	} else {
		for _, filePath := range args {
			isDir, _ := isDirectory(filePath)
			if !isDir {
				fmt.Printf("Skipped wrong path or having some issue %q:\n", filePath)
				continue
			}
			deleteDirectory(filePath)
		}
	}

	return nil
}

func runMKDIR(args []string, nestedDir bool) error {
	if !nestedDir {
		for _, dirName := range args {
			err := createSingleDir(dirName)
			if err != nil {
				fmt.Printf("Error creating directory %q:\n", dirName)
				continue
			}
		}
	} else {
		for _, dirName := range args {
			err := createNestedDir(dirName)
			if err != nil {
				fmt.Printf("Error creating nested directory %q:\n", dirName)
				continue
			}
		}
	}

	return nil
}
