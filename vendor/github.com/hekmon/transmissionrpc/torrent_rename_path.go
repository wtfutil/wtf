package transmissionrpc

import (
	"fmt"
)

/*
	Rename a torrent path
	https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L438
*/

// TorrentRenamePath allows to rename torrent name or path.
// 'path' is the path to the file or folder that will be renamed.
// 'name' the file or folder's new name
// https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L440
func (c *Client) TorrentRenamePath(id int64, path, name string) (err error) {
	if err = c.rpcCall("torrent-rename-path", torrentRenamePathPayload{
		IDs:  []int64{id},
		Path: path,
		Name: name,
	}, nil); err != nil {
		err = fmt.Errorf("'torrent-rename-path' rpc method failed: %v", err)
	}
	return
}

// TorrentRenamePathHash allows to rename torrent name or path by its hash.
// https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L440
func (c *Client) TorrentRenamePathHash(hash, path, name string) (err error) {
	if err = c.rpcCall("torrent-rename-path", torrentRenamePathHashPayload{
		Hashes: []string{hash},
		Path:   path,
		Name:   name,
	}, nil); err != nil {
		err = fmt.Errorf("'torrent-rename-path' rpc method failed: %v", err)
	}
	return
}

// https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L447
type torrentRenamePathPayload struct {
	IDs  []int64 `json:"ids"`  // the torrent torrent list, as described in 3.1 (must only be 1 torrent)
	Path string  `json:"path"` // the path to the file or folder that will be renamed
	Name string  `json:"name"` // the file or folder's new name
}

// https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L447
type torrentRenamePathHashPayload struct {
	Hashes []string `json:"ids"`  // the torrent torrent list, as described in 3.1 (must only be 1 torrent)
	Path   string   `json:"path"` // the path to the file or folder that will be renamed
	Name   string   `json:"name"` // the file or folder's new name
}
