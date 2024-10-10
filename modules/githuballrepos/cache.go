package githuballrepos

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"io/fs"

	"github.com/wtfutil/wtf/logger"
)

type Cacher interface {
	Get() *WidgetData
	Set(data *WidgetData)
	IsValid() bool
}

const cacheDuration = 5 * time.Minute

// Cache stores the widget data and its expiration time
type Cache struct {
	data       *WidgetData
	expires    time.Time
	configPath string
}

// NewCache creates a new Cache instance
func NewCache(configPath string) *Cache {
	cache := &Cache{
		configPath: configPath,
	}

	// Ensure the cache directory exists
	cacheDir := filepath.Dir(cache.configPath)
	logger.Log(fmt.Sprintf("Cache directory: %s\n", cacheDir))
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		logger.Log(fmt.Sprintf("Error creating cache directory: %s\n", err))
	}

	cache.load()
	return cache
}

// Set updates the cache with new data
func (c *Cache) Set(data *WidgetData) {
	c.data = data
	c.expires = time.Now().Add(cacheDuration)
	c.save()
}

// Get retrieves the cached data
func (c *Cache) Get() *WidgetData {
	return c.data
}

// IsValid checks if the cache is still valid
func (c *Cache) IsValid() bool {
	return c.data != nil && time.Now().Before(c.expires)
}

// save writes the cache data to disk
func (c *Cache) save() {
	cacheData := struct {
		Data    *WidgetData `json:"data"`
		Expires time.Time   `json:"expires"`
	}{
		Data:    c.data,
		Expires: c.expires,
	}

	jsonData, err := json.Marshal(cacheData)
	if err != nil {
		logger.Log(fmt.Sprintf("Error marshaling cache data: %s\n", err))
		return
	}

	if err := os.WriteFile(c.configPath, jsonData, fs.FileMode(0644)); err != nil {
		logger.Log(fmt.Sprintf("Error writing cache file: %s\n", err))
	}
}

// load reads the cache data from disk
func (c *Cache) load() {
	jsonData, err := os.ReadFile(c.configPath)
	if err != nil {
		if !os.IsNotExist(err) {
			logger.Log(fmt.Sprintf("Error reading cache file: %s\n", err))
		}
		return
	}

	var cacheData struct {
		Data    *WidgetData `json:"data"`
		Expires time.Time   `json:"expires"`
	}

	if err := json.Unmarshal(jsonData, &cacheData); err != nil {
		logger.Log(fmt.Sprintf("Error unmarshaling cache data: %s\n", err))
		return
	}

	c.data = cacheData.Data
	c.expires = cacheData.Expires
}
