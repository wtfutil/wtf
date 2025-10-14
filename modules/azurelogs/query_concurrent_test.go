package azurelogs

import (
	"fmt"
	"sync"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"
	"github.com/stretchr/testify/assert"
)

func TestLogQueryClients_ConcurrentAccess(t *testing.T) {
	// Save original state
	originalClients := LogQueryClients
	defer func() { LogQueryClients = originalClients }()

	// Reset to nil to test initialization
	LogQueryClients = nil

	const numGoroutines = 10
	const subscriptionID = "test-subscription"

	var wg sync.WaitGroup
	results := make([]bool, numGoroutines)

	// Launch multiple goroutines that try to access LogQueryClients simultaneously
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			// Use read lock to check if client exists
			clientsMutex.RLock()
			client := LogQueryClients[subscriptionID]
			clientsMapExists := LogQueryClients != nil
			clientsMutex.RUnlock()

			// Record if we found the map initialized
			results[index] = clientsMapExists

			// If map doesn't exist, try to initialize it
			if !clientsMapExists || client == nil {
				clientsMutex.Lock()
				// Double-check after acquiring write lock
				if LogQueryClients == nil {
					LogQueryClients = make(map[string]*azquery.LogsClient)
				}
				clientsMutex.Unlock()
			}
		}(i)
	}

	wg.Wait()

	// Verify that LogQueryClients was properly initialized
	assert.NotNil(t, LogQueryClients)
	assert.IsType(t, map[string]*azquery.LogsClient{}, LogQueryClients)

	// At least one goroutine should have seen the map as not existing initially
	anyFoundNil := false
	for _, result := range results {
		if !result {
			anyFoundNil = true
			break
		}
	}
	assert.True(t, anyFoundNil, "Expected at least one goroutine to see LogQueryClients as nil initially")
}

func TestLogQueryClients_ConcurrentReadWrite(t *testing.T) {
	// Save original state
	originalClients := LogQueryClients
	defer func() { LogQueryClients = originalClients }()

	// Initialize with a clean map
	LogQueryClients = make(map[string]*azquery.LogsClient)

	const numReaders = 5
	const numWriters = 3
	const subscriptionPrefix = "subscription-"

	var wg sync.WaitGroup

	// Launch reader goroutines
	for i := 0; i < numReaders; i++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()

			for j := 0; j < 10; j++ {
				// Read from the map safely
				clientsMutex.RLock()
				_ = LogQueryClients["test-subscription"]
				mapSize := len(LogQueryClients)
				clientsMutex.RUnlock()

				// Verify map size is reasonable (between 0 and numWriters)
				assert.GreaterOrEqual(t, mapSize, 0)
				assert.LessOrEqual(t, mapSize, numWriters)
			}
		}(i)
	}

	// Launch writer goroutines
	for i := 0; i < numWriters; i++ {
		wg.Add(1)
		go func(writerID int) {
			defer wg.Done()

			subscriptionID := subscriptionPrefix + string(rune('A'+writerID))

			// Write to the map safely
			clientsMutex.Lock()
			LogQueryClients[subscriptionID] = nil // Simulate adding a client entry
			clientsMutex.Unlock()
		}(i)
	}

	wg.Wait()

	// Verify final state
	clientsMutex.RLock()
	finalSize := len(LogQueryClients)
	clientsMutex.RUnlock()

	assert.Equal(t, numWriters, finalSize, "Expected exactly %d entries after concurrent writes", numWriters)
}

func TestLogQueryClients_RaceCondition(t *testing.T) {
	// This test verifies that concurrent access to LogQueryClients is thread-safe
	// Save original state
	originalClients := LogQueryClients
	defer func() { LogQueryClients = originalClients }()

	const numGoroutines = 100
	const numIterations = 10

	for attempt := 0; attempt < 5; attempt++ { // Run multiple attempts to catch race conditions
		// Reset to trigger concurrent initialization
		LogQueryClients = nil

		var wg sync.WaitGroup

		// Launch goroutines that all try to access/initialize LogQueryClients
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(goroutineID int) {
				defer wg.Done()

				subscriptionID := fmt.Sprintf("subscription-%d", goroutineID%3) // Use 3 different subscriptions

				for j := 0; j < numIterations; j++ {
					// Exactly match the logic from RunQuery function
					clientsMutex.RLock()
					client := LogQueryClients[subscriptionID]
					clientsMapExists := LogQueryClients != nil
					clientsMutex.RUnlock()

					if !clientsMapExists || client == nil {
						clientsMutex.Lock()
						// Double-check after acquiring write lock
						if LogQueryClients == nil {
							LogQueryClients = make(map[string]*azquery.LogsClient)
						}

						// Double-check for this specific subscription
						if LogQueryClients[subscriptionID] == nil {
							LogQueryClients[subscriptionID] = nil // Simulate client creation
						}
						clientsMutex.Unlock()
					}
				}
			}(i)
		}

		wg.Wait()

		// Verify final state is consistent - no race conditions occurred
		clientsMutex.RLock()
		assert.NotNil(t, LogQueryClients, "Attempt %d: LogQueryClients should be initialized", attempt+1)
		// Should have exactly 3 subscriptions (based on goroutineID%3)
		assert.Equal(t, 3, len(LogQueryClients), "Attempt %d: Expected 3 subscription entries", attempt+1)
		clientsMutex.RUnlock()
	}
}
