package hackernews

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func setupTestServer(handler http.HandlerFunc) (*httptest.Server, func()) {
	server := httptest.NewServer(handler)
	original := apiEndpoint
	apiEndpoint = server.URL + "/"
	return server, func() {
		apiEndpoint = original
		server.Close()
	}
}

func TestGetStories(t *testing.T) {
	tests := []struct {
		name       string
		storyType  string
		response   interface{}
		statusCode int
		wantIDs    []int
		wantErr    bool
		wantEmpty  bool
	}{
		{
			name:       "top stories",
			storyType:  "top",
			response:   []int{1, 2, 3, 4, 5},
			statusCode: http.StatusOK,
			wantIDs:    []int{1, 2, 3, 4, 5},
		},
		{
			name:       "new stories",
			storyType:  "new",
			response:   []int{10, 20, 30},
			statusCode: http.StatusOK,
			wantIDs:    []int{10, 20, 30},
		},
		{
			name:       "job stories",
			storyType:  "job",
			response:   []int{100},
			statusCode: http.StatusOK,
			wantIDs:    []int{100},
		},
		{
			name:       "ask stories",
			storyType:  "ask",
			response:   []int{50, 51},
			statusCode: http.StatusOK,
			wantIDs:    []int{50, 51},
		},
		{
			name:       "case insensitive - TOP",
			storyType:  "TOP",
			response:   []int{7, 8},
			statusCode: http.StatusOK,
			wantIDs:    []int{7, 8},
		},
		{
			name:       "unknown story type returns empty",
			storyType:  "unknown",
			response:   nil,
			statusCode: http.StatusOK,
			wantEmpty:  true,
		},
		{
			name:       "server error",
			storyType:  "top",
			response:   nil,
			statusCode: http.StatusInternalServerError,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
				if tt.statusCode != http.StatusOK {
					w.WriteHeader(tt.statusCode)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(tt.response)
			})
			defer cleanup()
			_ = server

			ids, err := GetStories(tt.storyType)

			if tt.wantErr {
				assert.Assert(t, err != nil)
				return
			}

			assert.NilError(t, err)

			if tt.wantEmpty {
				assert.Equal(t, 0, len(ids))
				return
			}

			assert.DeepEqual(t, tt.wantIDs, ids)
		})
	}
}

func TestGetStories_InvalidJSON(t *testing.T) {
	server, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("not valid json"))
	})
	defer cleanup()
	_ = server

	_, err := GetStories("top")
	assert.Assert(t, err != nil)
}

func TestGetStory(t *testing.T) {
	tests := []struct {
		name       string
		id         int
		response   Story
		statusCode int
		wantErr    bool
	}{
		{
			name: "valid story",
			id:   123,
			response: Story{
				By:          "author1",
				Descendants: 42,
				ID:          123,
				Kids:        []int{456, 789},
				Score:       100,
				Time:        1617000000,
				Title:       "Test Story Title",
				Type:        "story",
				URL:         "https://example.com/article",
			},
			statusCode: http.StatusOK,
		},
		{
			name: "story without URL (Ask HN)",
			id:   456,
			response: Story{
				By:    "asker",
				ID:    456,
				Score: 50,
				Title: "Ask HN: Something?",
				Type:  "story",
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "server error",
			id:         999,
			statusCode: http.StatusServiceUnavailable,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
				if tt.statusCode != http.StatusOK {
					w.WriteHeader(tt.statusCode)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(tt.response)
			})
			defer cleanup()
			_ = server

			story, err := GetStory(tt.id)

			if tt.wantErr {
				assert.Assert(t, err != nil)
				return
			}

			assert.NilError(t, err)
			assert.Equal(t, tt.response.ID, story.ID)
			assert.Equal(t, tt.response.Title, story.Title)
			assert.Equal(t, tt.response.By, story.By)
			assert.Equal(t, tt.response.URL, story.URL)
			assert.Equal(t, tt.response.Score, story.Score)
			assert.Equal(t, tt.response.Type, story.Type)
		})
	}
}

func TestGetStory_InvalidJSON(t *testing.T) {
	server, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("{invalid"))
	})
	defer cleanup()
	_ = server

	_, err := GetStory(1)
	assert.Assert(t, err != nil)
}

func TestApiRequest_RequestPath(t *testing.T) {
	var receivedPath string
	server, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		receivedPath = r.URL.Path
		_, _ = w.Write([]byte("[]"))
	})
	defer cleanup()
	_ = server

	_, err := apiRequest("topstories")
	assert.NilError(t, err)
	assert.Equal(t, "/topstories.json", receivedPath)
}

func TestApiRequest_ItemPath(t *testing.T) {
	var receivedPath string
	server, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		receivedPath = r.URL.Path
		_, _ = w.Write([]byte("{}"))
	})
	defer cleanup()
	_ = server

	_, err := apiRequest("item/123")
	assert.NilError(t, err)
	assert.Equal(t, "/item/123.json", receivedPath)
}

func TestApiRequest_HTTPMethod(t *testing.T) {
	var receivedMethod string
	server, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		receivedMethod = r.Method
		_, _ = w.Write([]byte("{}"))
	})
	defer cleanup()
	_ = server

	_, err := apiRequest("topstories")
	assert.NilError(t, err)
	assert.Equal(t, "GET", receivedMethod)
}

func TestApiRequest_Various4xxErrors(t *testing.T) {
	codes := []int{
		http.StatusBadRequest,
		http.StatusNotFound,
		http.StatusForbidden,
	}

	for _, code := range codes {
		t.Run(http.StatusText(code), func(t *testing.T) {
			server, cleanup := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(code)
			})
			defer cleanup()
			_ = server

			_, err := apiRequest("topstories")
			assert.Assert(t, err != nil)
		})
	}
}
