package transmissionrpc

import (
	"errors"
	"fmt"
)

/*
	Removing a Torrent
	https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L407
*/

// TorrentRemove allows to delete one or more torrents only or with their data.
func (c *Client) TorrentRemove(payload *TorrentRemovePayload) (err error) {
	// Validate
	if payload == nil {
		return errors.New("payload can't be nil")
	}
	// Send payload
	if err = c.rpcCall("torrent-remove", payload, nil); err != nil {
		return fmt.Errorf("'torrent-remove' rpc method failed: %v", err)
	}
	return
}

// TorrentRemovePayload holds the torrent id(s) to delete with a data deletion flag.
// https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L413
type TorrentRemovePayload struct {
	IDs             []int64 `json:"ids"`
	DeleteLocalData bool    `json:"delete-local-data"`
}
