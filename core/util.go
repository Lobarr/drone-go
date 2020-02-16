package core

import (
	"os"
)

//fileExists checks if a file exists
func fileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

//isFile checks if path is a file
func isFile(filePath string) bool {
	fileInfo, _ := os.Stat(filePath)
	return fileInfo.Mode().IsRegular()
}

//getFileSize gets file size
func getFileSize(filePath string) int64 {
	fileInfo, _ := os.Stat(filePath)
	return fileInfo.Size()
}
