package cfg

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/olebedev/config"
)

const (
	// XdgConfigDir defines the path to the minimal XDG-compatible configuration directory
	XdgConfigDir = "~/.config/"

	// WtfConfigDirV1 defines the path to the first version of configuration. Do not use this
	WtfConfigDirV1 = "~/.wtf/"

	// WtfConfigDirV2 defines the path to the second version of the configuration. Use this.
	WtfConfigDirV2 = "~/.config/wtf/"

	// WtfConfigFile defines the name of the default config file
	WtfConfigFile = "config.yml"
)

/* -------------------- Exported Functions -------------------- */

// CreateFile creates the named file in the config directory, if it does not already exist.
// If the file exists it does not recreate it.
// If successful, returns the absolute path to the file
// If unsuccessful, returns an error
func CreateFile(fileName string) (string, error) {
	configDir, err := WtfConfigDir()
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(configDir, fileName)

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

// Initialize takes care of settings up the initial state of WTF configuration
// It ensures necessary directories and files exist
func Initialize(hasCustom bool) {
	if !hasCustom {
		migrateOldConfig()
	}

	// These always get created because this is where modules should write any permanent
	// data they need to persist between runs (i.e.: log, textfile, etc.)
	createWtfConfigDir()

	if !hasCustom {
		createWtfConfigFile()
		chmodConfigFile()
	}
}

// WtfConfigDir returns the absolute path to the configuration directory
func WtfConfigDir() (string, error) {
	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		configDir = WtfConfigDirV2
	} else {
		configDir += "/wtf/"
	}
	configDir, err := expandHomeDir(configDir)
	if err != nil {
		return "", err
	}

	return configDir, nil
}

// LoadWtfConfigFile loads the specified config file
func LoadWtfConfigFile(filePath string) *config.Config {
	absPath, _ := expandHomeDir(filePath)

	cfg, err := config.ParseYamlFile(absPath)
	if err != nil {
		displayWtfConfigFileLoadError(absPath, err)
		os.Exit(1)
	}

	return cfg
}

/* -------------------- Unexported Functions -------------------- */

// chmodConfigFile sets the mode of the config file to r+w for the owner only
func chmodConfigFile() {
	configDir, _ := WtfConfigDir()
	relPath := filepath.Join(configDir, WtfConfigFile)
	absPath, _ := expandHomeDir(relPath)

	_, err := os.Stat(absPath)
	if err != nil && os.IsNotExist(err) {
		return
	}

	err = os.Chmod(absPath, 0600)
	if err != nil {
		return
	}
}

// createWtfConfigDir creates the necessary directories for storing the default config file
// If ~/.config/wtf is missing, it will try to create it
func createWtfConfigDir() {
	wtfConfigDir, _ := WtfConfigDir()

	if _, err := os.Stat(wtfConfigDir); os.IsNotExist(err) {
		err := os.MkdirAll(wtfConfigDir, os.ModePerm)
		if err != nil {
			displayWtfConfigDirCreateError(err)
			os.Exit(1)
		}
	}
}

// createWtfConfigFile creates a simple config file in the config directory if
// one does not already exist
func createWtfConfigFile() {
	filePath, err := CreateFile(WtfConfigFile)
	if err != nil {
		displayDefaultConfigCreateError(err)
		os.Exit(1)
	}

	// If the file is empty, write to it
	file, _ := os.Stat(filePath)

	if file.Size() == 0 {
		if os.WriteFile(filePath, []byte(defaultConfigFile), 0600) != nil {
			displayDefaultConfigWriteError(err)
			os.Exit(1)
		}
	}
}

// Expand expands the path to include the home directory if the path
// is prefixed with `~`. If it isn't prefixed with `~`, the path is
// returned as-is.
func expandHomeDir(path string) (string, error) {
	if path == "" {
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

// migrateOldConfig copies any existing configuration from the old location
// to the new, XDG-compatible location
func migrateOldConfig() {
	srcDir, _ := expandHomeDir(WtfConfigDirV1)
	destDir, _ := WtfConfigDir()

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
