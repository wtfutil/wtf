package logger

import (
	"log"
	"os"
	"path/filepath"
)

/* -------------------- Exported Functions -------------------- */

func Log(msg string) {
	if LogFileMissing() {
		return
	}

	f, err := os.OpenFile(LogFilePath(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer func() { _ = f.Close() }()

	log.SetOutput(f)
	log.Println(msg)
}

func LogFileMissing() bool {
	return LogFilePath() == ""
}

func LogFilePath() string {
	// Check WTF_CONFIG_DIR first for hermetic test environments
	if configDir := os.Getenv("WTF_CONFIG_DIR"); configDir != "" {
		return filepath.Join(configDir, "log.txt")
	}

	dir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	return filepath.Join(dir, ".config", "wtf", "log.txt")
}
