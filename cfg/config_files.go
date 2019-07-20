package cfg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/olebedev/config"
)

// XdgConfigDir defines the path to the minimal XDG-compatible configuration directory
const XdgConfigDir = "~/.config/"

// WtfConfigDirV1 defines the path to the first version of configuration. Do not use this
const WtfConfigDirV1 = "~/.wtf/"

// WtfConfigDirV2 defines the path to the second version of the configuration. Use this.
const WtfConfigDirV2 = "~/.config/wtf/"

/* -------------------- Config Migration -------------------- */

// MigrateOldConfig copies any existing configuration from the old location
// to the new, XDG-compatible location
func MigrateOldConfig() {
	srcDir, _ := expandHomeDir(WtfConfigDirV1)
	destDir, _ := expandHomeDir(WtfConfigDirV2)

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
	}

	// Delete the old directory if the new one exists
	if _, err := os.Stat(destDir); err == nil {
		err := os.RemoveAll(srcDir)
		if err != nil {
			fmt.Println(err)
		}
	}
}

/* -------------------- Config Migration -------------------- */

// ConfigDir returns the absolute path to the configuration directory
func WtfConfigDir() (string, error) {
	configDir, err := expandHomeDir(WtfConfigDirV2)
	if err != nil {
		return "", err
	}

	return configDir, nil
}

// CreateXdgConfigDir creates the necessary base directory for storing the config file
// If ~/.config is missing, it will try to create it
func CreateXdgConfigDir() {
	xdgConfigDir, _ := expandHomeDir(XdgConfigDir)

	if _, err := os.Stat(xdgConfigDir); os.IsNotExist(err) {
		err := os.Mkdir(xdgConfigDir, os.ModePerm)
		if err != nil {
			displayXdgConfigDirCreateError(err)
			os.Exit(1)
		}
	}
}

// CreateWtfConfigDir creates the necessary directories for storing the default config file
// If ~/.config/wtf is missing, it will try to create it
func CreateWtfConfigDir() {
	wtfConfigDir, _ := WtfConfigDir()

	if _, err := os.Stat(wtfConfigDir); os.IsNotExist(err) {
		err := os.Mkdir(wtfConfigDir, os.ModePerm)
		if err != nil {
			displayWtfConfigDirCreateError(err)
			os.Exit(1)
		}
	}
}

// CreateWtfConfigFile creates a simple config file in the config directory if
// one does not already exist
func CreateWtfConfigFile() {
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
	configDir, err := WtfConfigDir()
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

// LoadWtfConfigFile loads the config.yml file to configure the app
func LoadWtfConfigFile(filePath string, isCustomConfig bool) *config.Config {
	absPath, _ := expandHomeDir(filePath)

	cfg, err := config.ParseYamlFile(absPath)
	if err != nil {
		if isCustomConfig {
			displayWtfCustomConfigFileLoadError(err)
		} else {
			displayWtfConfigFileLoadError(err)
		}

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

/* -------------------- Unexported Functions -------------------- */

func displayXdgConfigDirCreateError(err error) {
	fmt.Printf("\n\033[1mERROR:\033[0m Could not create the '\033[0;33m%s\033[0m' directory.\n", XdgConfigDir)
	fmt.Println()
	fmt.Printf("Error: \033[0;31m%s\033[0m\n\n", err.Error())
}

func displayWtfConfigDirCreateError(err error) {
	fmt.Printf("\n\033[1mERROR:\033[0m Could not create the '\033[0;33m%s\033[0m' directory.\n", WtfConfigDirV2)
	fmt.Println()
	fmt.Printf("Error: \033[0;31m%s\033[0m\n\n", err.Error())
}

func displayWtfConfigFileLoadError(err error) {
	fmt.Println("\n\033[1mERROR:\033[0m Could not load '\033[0;33mconfig.yml\033[0m'.")
	fmt.Println()
	fmt.Println("This could mean one of two things:")
	fmt.Println()
	fmt.Println("    1. Your \033[0;33mconfig.yml\033[0m file is missing. Check in \033[0;33m~/.config/wtf\033[0m to see if \033[0;33mconfig.yml\033[0m is there.")
	fmt.Println("    2. Your \033[0;33mconfig.yml\033[0m file has a syntax error. Try running it through http://www.yamllint.com to check for errors.")
	fmt.Println()
	fmt.Printf("Error: \033[0;31m%s\033[0m\n\n", err.Error())
}

func displayWtfCustomConfigFileLoadError(err error) {
	fmt.Println("\n\033[1mERROR:\033[0m Could not load '\033[0;33mconfig.yml\033[0m'.")
	fmt.Println()
	fmt.Println("This could mean one of two things:")
	fmt.Println()
	fmt.Println("    1. That file doesn't exist.")
	fmt.Println("    2. That file has a YAML syntax error. Try running it through http://www.yamllint.com to check for errors.")
	fmt.Println()
	fmt.Printf("Error: \033[0;31m%s\033[0m\n\n", err.Error())
}

// Expand expands the path to include the home directory if the path
// is prefixed with `~`. If it isn't prefixed with `~`, the path is
// returned as-is.
func expandHomeDir(path string) (string, error) {
	if len(path) == 0 {
		return path, nil
	}

	if path[0] != '~' {
		return path, nil
	}

	if len(path) > 1 && path[1] != '/' && path[1] != '\\' {
		return "", errors.New("cannot expand user-specific home dir")
	}

	dir, err := home()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, path[1:]), nil
}

// Dir returns the home directory for the executing user.
// An error is returned if a home directory cannot be detected.
func home() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	if currentUser.HomeDir == "" {
		return "", errors.New("cannot find user-specific home dir")
	}

	return currentUser.HomeDir, nil
}
