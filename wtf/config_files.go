package wtf

import (
	"fmt"
	"os"

	"github.com/olebedev/config"
)

func ConfigDir() (string, error) {
	configDir, err := ExpandHomeDir("~/.wtf/")
	if err != nil {
		return "", err
	}

	return configDir, nil
}

// CreateConfigDir creates the .wtf directory in the user's home dir
func CreateConfigDir() bool {
	configDir, _ := ConfigDir()

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.Mkdir(configDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	return true
}

// CreateFile creates the named file in the config directory, if it does not already exist.
// If the file exists it does not recreate it.
// If successful, eturns the absolute path to the file
// If unsuccessful, returns an error
func CreateFile(fileName string) (string, error) {
	configDir, err := ConfigDir()
	if err != nil {
		return "", err
	}

	filePath := fmt.Sprintf("%s/%s", configDir, fileName)

	// Check if the file already exists; if it does not, create it
	_, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(filePath)
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}

	return filePath, nil
}

// LoadConfigFile loads the config.yml file to configure the app
func LoadConfigFile(filePath string) *config.Config {
	absPath, _ := ExpandHomeDir(filePath)

	cfg, err := config.ParseYamlFile(absPath)
	if err != nil {
		fmt.Println("\n\n\033[1m ERROR:\033[0m Could not load '\033[0;33mconfig.yml\033[0m'.\n Please add a \033[0;33mconfig.yml\033[0m file to your \033[0;33m~/.wtf\033[0m directory.\n See \033[1;34mhttps://github.com/senorprogrammer/wtf\033[0m for details.\n")
		fmt.Printf(" %s\n", err.Error())
		os.Exit(1)
	}

	return cfg
}

func ReadFile(fileName string) (string, error) {
	configDir, err := ConfigDir()
	if err != nil {
		return "", err
	}

	filePath := fmt.Sprintf("%s/%s", configDir, fileName)

	fileData, err := ReadFileBytes(filePath)
	if err != nil {
		return "", err
	}

	return string(fileData), nil
}
