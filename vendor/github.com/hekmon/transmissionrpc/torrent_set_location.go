package transmissionrpc

import (
	"fmt"
)

/*
	Moving a torrent
	https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L421
*/

// TorrentSetLocation allows to set a new location for one or more torrents.
// 'location' is the new torrent location.
// 'move' if true, move from previous location. Otherwise, search "location" for file.
// https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L423
func (c *Client) TorrentSetLocation(id int64, location string, move bool) (err error) {
	if err = c.rpcCall("torrent-set-location", torrentSetLocationPayload{
		IDs:      []int64{id},
		Location: location,
		Move:     move,
	}, nil); err != nil {
		err = fmt.Errorf("'torrent-set-location' rpc method failed: %v", err)
	}
	return
}

// TorrentSetLocationHash allows to set a new location for one or more torrents.
// 'location' is the new torrent location.
// 'move' if true, move from previous location. Otherwise, search "location" for file.
// https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L423
func (c *Client) TorrentSetLocationHash(hash, location string, move bool) (err error) {
	if err = c.rpcCall("torrent-set-location", torrentSetLocationHashPayload{
		Hashes:   []string{hash},
		Location: location,
		Move:     move,
	}, nil); err != nil {
		err = fmt.Errorf("'torrent-set-location' rpc method failed: %v", err)
	}
	return
}

// https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L427
type torrentSetLocationPayload struct {
	IDs      []int64 `json:"ids"`      // torrent list
	Location string  `json:"location"` // the new torrent location
	Move     bool    `json:"move"`     // if true, move from previous location. Otherwise, search "location" for files
}

// https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L427
type torrentSetLocationHashPayload struct {
	Hashes   []string `json:"ids"`      // torrent list
	Location string   `json:"location"` // the new torrent location
	Move     bool     `json:"move"`     // if true, move from previous location. Otherwise, search "location" for files
}
