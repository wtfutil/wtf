package logger

import (
	"log"
	"os"
	"path/filepath"

	"github.com/wtfutil/wtf/utils"
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
	defer f.Close()

	log.SetOutput(f)
	log.Println(msg)
}

func LogFileMissing() bool {
	return LogFilePath() == ""
}

func LogFilePath() string {
	dir, err := utils.Home()
	if err != nil {
		return ""
	}

	return filepath.Join(dir, ".config", "wtf", "log.txt")
}
