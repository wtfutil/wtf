package cfg

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/logger"
	"github.com/senorprogrammer/wtf/wtf"
)

const CONFIG_DIR_V1 = "~/.wtf/"
const CONFIG_DIR_V2 = "~/.config/wtf/"

/* -------------------- Config Migration -------------------- */

// MigrateOldConfig copies any existing configuration from the old location
// to the new, XDG-compatible location
func MigrateOldConfig() {
	srcDir, _ := wtf.ExpandHomeDir(CONFIG_DIR_V1)
	destDir, _ := wtf.ExpandHomeDir(CONFIG_DIR_V2)

	// If the old config directory doesn't exist, do not move
	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		return
	}

	// If the new config directory already exists, do not move
	if _, err := os.Stat(destDir); err == nil {
		return
	}

	// Time to move
	err := Copy(srcDir, destDir)
	if err != nil {
		panic(err)
	} else {
		logger.Log(fmt.Sprintf("Copied old config from %s to %s", srcDir, destDir))
	}

	// Delete the old directory if the new one exists
	if _, err := os.Stat(destDir); err == nil {
		err := os.RemoveAll(srcDir)
		if err != nil {
			logger.Log(err.Error())
		}
	}
}

/* -------------------- Config Migration -------------------- */

// ConfigDir returns the absolute path to the configuration directory
func ConfigDir() (string, error) {
	configDir, err := wtf.ExpandHomeDir(CONFIG_DIR_V2)
	if err != nil {
		return "", err
	}

	return configDir, nil
}

// CreateConfigDir creates the wtf/ directory in the user's home dir
func CreateConfigDir() {
	configDir, _ := ConfigDir()

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.Mkdir(configDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

// CreateConfigFile creates a simple config file in the config directory if
// one does not already exist
func CreateConfigFile() {
	filePath, err := CreateFile("config.yml")
	if err != nil {
		panic(err)
	}

	// If the file is empty, write to it
	file, _ := os.Stat(filePath)

	if file.Size() == 0 {
		if ioutil.WriteFile(filePath, []byte(simpleConfig), 0644) != nil {
			panic(err)
		}
	}
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
	absPath, _ := wtf.ExpandHomeDir(filePath)

	cfg, err := config.ParseYamlFile(absPath)
	if err != nil {
		fmt.Println("\n\n\033[1m ERROR:\033[0m Could not load '\033[0;33mconfig.yml\033[0m'.\n Please add a \033[0;33mconfig.yml\033[0m file to your \033[0;33m~/.wtf\033[0m directory.\n See \033[1;34mhttps://github.com/senorprogrammer/wtf\033[0m for details.")
		fmt.Printf(" %s\n", err.Error())
		os.Exit(1)
	}

	return cfg
}

const simpleConfig = `wtf:
  colors:
    border:
      focusable: darkslateblue
      focused: orange
      normal: gray
  grid:
    columns: [40, 40]
    rows: [13, 13, 4]
  refreshInterval: 1
  mods:
    clocks:
      colors:
        rows:
          even: "lightblue"
          odd: "white"
      enabled: true
      locations:
        Avignon: "Europe/Paris"
        Barcelona: "Europe/Madrid"
        Dubai: "Asia/Dubai"
        Vancouver: "America/Vancouver"
        Toronto: "America/Toronto"
      position:
        top: 0
        left: 0
        height: 1
        width: 1
      refreshInterval: 15
      sort: "alphabetical"
    security:
      enabled: true
      position:
        top: 1
        left: 0
        height: 1
        width: 1
      refreshInterval: 3600
    status:
      enabled: true
      position:
        top: 2
        left: 0
        height: 1
        width: 2
      refreshInterval: 1
    system:
      enabled: true
      position:
        top: 0
        left: 1
        height: 1
        width: 1
      refreshInterval: 3600
    textfile:
      enabled: true
      filePath: "~/.config/wtf/config.yml"
      position:
        top: 1
        left: 1
        height: 1
        width: 1
      refreshInterval: 30
`
