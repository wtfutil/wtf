package wtf

import (
	"fmt"
	"os"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/homedir"
)

// CreateConfigDir creates the .wtf directory in the user's home dir
func CreateConfigDir() bool {
	homeDir, _ := homedir.Expand("~/.wtf/")

	if _, err := os.Stat(homeDir); os.IsNotExist(err) {
		err := os.Mkdir(homeDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	return true
}

// LoadConfigFile loads the config.yml file to configure the app
func LoadConfigFile(filePath string) *config.Config {
	absPath, _ := homedir.Expand(filePath)

	cfg, err := config.ParseYamlFile(absPath)
	if err != nil {
		fmt.Println("\n\n\033[1m ERROR:\033[0m Could not load '\033[0;33mconfig.yml\033[0m'.\n Please add a \033[0;33mconfig.yml\033[0m file to your \033[0;33m~/.wtf\033[0m directory.\n See \033[1;34mhttps://github.com/senorprogrammer/wtf\033[0m for details.\n\n")
		os.Exit(1)
	}

	return cfg
}
