package config

import (
	"log"
	"os"
)

var (
	StoragePath string
	UploadsPath string
	LogsPath    string
	TmpPath     string
)

func InitStorage() {
	StoragePath = os.Getenv("STORAGE_PATH")
	if StoragePath == "" {
		StoragePath = "storage"
	}

	UploadsPath = os.Getenv("UPLOADS_PATH")
	if UploadsPath == "" {
		UploadsPath = StoragePath + "/uploads"
	}

	LogsPath = os.Getenv("LOGS_PATH")
	if LogsPath == "" {
		LogsPath = StoragePath + "/logs"
	}

	TmpPath = os.Getenv("TMP_PATH")
	if TmpPath == "" {
		TmpPath = StoragePath + "/tmp"
	}

	dirs := []string{StoragePath, UploadsPath, LogsPath, TmpPath}
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				log.Fatalf("❌ Could not create directory %s: %v", dir, err)
			}
			log.Printf("✅ Created directory: %s", dir)
		}
	}
}
