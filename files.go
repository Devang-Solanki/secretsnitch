/*
*
Worker-optimized downloading
Stress tested with 100k URLs from GitHub - took around 2 minutes on an i5-8350U / 16GB DDR3 @ 1333 MHz /
*
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	cacheDir           = ".urlCache/"
	cacheFileExtension = ".cache"
	defaultOutputDir   = "output.json"
)

func listFiles(directory string) ([]string, error) {
	var files []string

	dir, err := os.Open(directory)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	fileInfo, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	for _, file := range fileInfo {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}

	return files, nil
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		host := scanner.Text()
		lines = append(lines, host)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func makeDir(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	return nil
}

func fileExists(location string) bool {
	if _, err := os.Stat(location); err == nil {
		return true
	}
	return false
}

func makeCacheFilename(url string) string {
	return cacheDir + md5Hash(url)[0:8] + cacheFileExtension
}

func listCachedFiles() ([]string, error) {
	var fileList []string
	files, err := os.ReadDir(cacheDir)
	if err != nil {
		log.Fatal(err)
		return fileList, err
	}
	for _, file := range files {
		if !file.IsDir() {
			relativePath := cacheDir + file.Name()
			fileList = append(fileList, relativePath)
		}
	}
	return fileList, err
}

func appendToFile(filePath string, text string) error {
	dir := filepath.Dir(filePath)

	if dir != "." {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directories: %w", err)
		}
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	if _, err := file.WriteString(text + "\n"); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func logSecret(secret ToolData, outputFile *string) {
	unindented, _ := json.Marshal(secret)
	// indented, _ := json.MarshalIndent(secrets, "", "	")
	appendToFile(*outputFile, string(unindented))
}
