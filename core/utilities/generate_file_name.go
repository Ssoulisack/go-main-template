package utilities

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func GenerateNewFileName(cadId string) string {
	// You can use any method to generate a new file name here
	// For example, appending a timestamp to the old name
	id := uuid.New().String()
	timestamp := time.Now().Unix()
	newName := fmt.Sprintf("%d_%v_%s.png", timestamp, id, cadId)
	return newName
}

func GetFileNameInURL(url string) string {
	if url == "" {
		return ""
	}
	fileName := filepath.Base(url)
	return fileName
}
