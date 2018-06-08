package wtf

import (
	"log"
	"os"
	"path/filepath"
)

//Log basic message logging, defaults to ~/.wtf/log.txt
func Log(message string) {

	dir, err := Home()
	if err != nil {
		return
	}

	logfile := filepath.Join(dir, ".wtf", "log.txt")
	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(message)
}
