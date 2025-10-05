package jira

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestUserIDCacheMap_SetAndGet(t *testing.T) {
	cache := &UserIDCacheMap{
		cache: make(map[string]UserIDCache),
	}

	// Test setting and getting a value
	username := "testuser"
	accountID := "account:123456789"
	duration := 5 * time.Minute

	cache.Set(username, accountID, duration)

	// Test successful retrieval
	retrievedID, found := cache.Get(username)
	assert.Equal(t, true, found)
	assert.Equal(t, accountID, retrievedID)
}

func TestUserIDCacheMap_GetNonExistent(t *testing.T) {
	cache := &UserIDCacheMap{
		cache: make(map[string]UserIDCache),
	}

	// Test getting non-existent value
	retrievedID, found := cache.Get("nonexistent")
	assert.Equal(t, false, found)
	assert.Equal(t, "", retrievedID)
}

func TestUserIDCacheMap_GetExpired(t *testing.T) {
	cache := &UserIDCacheMap{
		cache: make(map[string]UserIDCache),
	}

	// Set an entry that expires immediately
	username := "expireduser"
	accountID := "account:987654321"
	cache.Set(username, accountID, -1*time.Second) // Already expired

	// Test that expired entry is not returned and is cleaned up
	retrievedID, found := cache.Get(username)
	assert.Equal(t, false, found)
	assert.Equal(t, "", retrievedID)

	// Verify the expired entry was removed from cache
	cache.mutex.RLock()
	_, exists := cache.cache[username]
	cache.mutex.RUnlock()
	assert.Equal(t, false, exists)
}

func TestUserIDCacheMap_Clear(t *testing.T) {
	cache := &UserIDCacheMap{
		cache: make(map[string]UserIDCache),
	}

	// Add a valid entry and an expired entry
	cache.Set("validuser", "account:111", 5*time.Minute)
	cache.Set("expireduser", "account:222", -1*time.Second)

	// Clear expired entries
	cache.Clear()

	// Valid entry should still exist
	_, found := cache.Get("validuser")
	assert.Equal(t, true, found)

	// Expired entry should be gone
	cache.mutex.RLock()
	_, exists := cache.cache["expireduser"]
	cache.mutex.RUnlock()
	assert.Equal(t, false, exists)
}

func TestExtractAccountIDFromJQL(t *testing.T) {
	tests := []struct {
		name     string
		jql      string
		expected string
	}{
		{
			name:     "valid account ID",
			jql:      `assignee = "account:5b10ac8d82e05b22cc7d4ef5"`,
			expected: "account:5b10ac8d82e05b22cc7d4ef5",
		},
		{
			name:     "single quotes",
			jql:      `assignee = 'account:123456789'`,
			expected: "", // Our function only handles double quotes
		},
		{
			name:     "no quotes",
			jql:      `assignee = account:123456789`,
			expected: "",
		},
		{
			name:     "empty string",
			jql:      "",
			expected: "",
		},
		{
			name:     "malformed JQL",
			jql:      `assignee = "incomplete`,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractAccountIDFromJQL(tt.jql)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestConvertJQLWithUsername_CacheHit(t *testing.T) {
	// Setup a mock widget (minimal setup for testing)
	widget := &Widget{}

	// Clear and setup cache
	userIDCache = &UserIDCacheMap{
		cache: make(map[string]UserIDCache),
	}

	// Pre-populate cache
	username := "cacheduser"
	accountID := "account:cached123"
	userIDCache.Set(username, accountID, 5*time.Minute)

	// Test that cached value is returned without API call
	result, err := widget.ConvertJQLWithUsername(username)

	assert.NilError(t, err)
	assert.Equal(t, `assignee = "account:cached123"`, result)
}

func TestConvertJQLWithUsername_APICalls(t *testing.T) {
	// Create a mock JIRA server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify it's a POST request to the right endpoint
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/rest/api/3/jql/pdcleaner", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		// Mock response
		response := JQLConversionResponse{
			QueryStrings: []ConvertedQuery{
				{
					Query:          `assignee = "testuser"`,
					ConvertedQuery: `assignee = "account:5b10ac8d82e05b22cc7d4ef5"`,
					UserMessages:   []UserMessage{},
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Setup widget with mock server
	widget := &Widget{
		settings: &Settings{
			domain: server.URL,
		},
	}

	// Clear cache
	userIDCache = &UserIDCacheMap{
		cache: make(map[string]UserIDCache),
	}

	// Test API call
	result, err := widget.ConvertJQLWithUsername("testuser")

	assert.NilError(t, err)
	assert.Equal(t, `assignee = "account:5b10ac8d82e05b22cc7d4ef5"`, result)

	// Verify it was cached
	cachedID, found := userIDCache.Get("testuser")
	assert.Equal(t, true, found)
	assert.Equal(t, "account:5b10ac8d82e05b22cc7d4ef5", cachedID)
}

func TestConvertJQLWithUsername_APIError(t *testing.T) {
	// Create a mock server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Internal Server Error"))
	}))
	defer server.Close()

	// Setup widget with mock server
	widget := &Widget{
		settings: &Settings{
			domain: server.URL,
		},
	}

	// Clear cache
	userIDCache = &UserIDCacheMap{
		cache: make(map[string]UserIDCache),
	}

	// Test API error handling
	result, err := widget.ConvertJQLWithUsername("testuser")

	assert.ErrorContains(t, err, "500 Internal Server Error")
	assert.Equal(t, "", result)
}

func TestConvertJQLWithUsername_EmptyResponse(t *testing.T) {
	// Create a mock server that returns empty response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := JQLConversionResponse{
			QueryStrings: []ConvertedQuery{},
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Setup widget with mock server
	widget := &Widget{
		settings: &Settings{
			domain: server.URL,
		},
	}

	// Clear cache
	userIDCache = &UserIDCacheMap{
		cache: make(map[string]UserIDCache),
	}

	// Test empty response handling
	result, err := widget.ConvertJQLWithUsername("testuser")

	assert.Error(t, err, "no conversion result for username: testuser")
	assert.Equal(t, "", result)
}

func TestConvertJQLWithUsername_InvalidAccountID(t *testing.T) {
	// Create a mock server that returns malformed JQL
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := JQLConversionResponse{
			QueryStrings: []ConvertedQuery{
				{
					Query:          `assignee = "testuser"`,
					ConvertedQuery: `assignee = malformed_without_quotes`,
					UserMessages:   []UserMessage{},
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Setup widget with mock server
	widget := &Widget{
		settings: &Settings{
			domain: server.URL,
		},
	}

	// Clear cache
	userIDCache = &UserIDCacheMap{
		cache: make(map[string]UserIDCache),
	}

	// Test invalid account ID handling
	result, err := widget.ConvertJQLWithUsername("testuser")

	assert.ErrorContains(t, err, "failed to extract account ID from converted query")
	assert.Equal(t, "", result)
}
