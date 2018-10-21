package spotify

// This file contains the types that implement Spotify's cursor-based
// paging object.  Like the standard paging object, this object is a
// container for a set of items. Unlike the standard paging object, a
// cursor-based paging object does not provide random access to the results.

// Cursor contains a key that can be used to find the next set
// of items.
type Cursor struct {
	After string `json:"after"`
}

// cursorPage contains all of the fields in a Spotify cursor-based
// paging object, except for the actual items.  This type is meant
// to be embedded in other types that add the Items field.
type cursorPage struct {
	// A link to the Web API endpoint returning the full
	// result of this request.
	Endpoint string `json:"href"`
	// The maximum number of items returned, as set in the query
	// (or default value if unset).
	Limit int `json:"limit"`
	// The URL to the next set of items.
	Next string `json:"next"`
	// The total number of items available to return.
	Total int `json:"total"`
	// The cursor used to find the next set of items.
	Cursor Cursor `json:"cursors"`
}

// FullArtistCursorPage is a cursor-based paging object containing
// a set of FullArtist objects.
type FullArtistCursorPage struct {
	cursorPage
	Artists []FullArtist `json:"items"`
}
