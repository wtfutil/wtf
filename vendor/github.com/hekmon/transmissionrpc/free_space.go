package transmissionrpc

import (
	"fmt"

	"github.com/hekmon/cunits"
)

/*
	Free Space
	https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L618
*/

// FreeSpace allow to see how much free space is available in a client-specified folder.
func (c *Client) FreeSpace(path string) (freeSpace cunits.Bits, err error) {
	payload := &transmissionFreeSpacePayload{Path: path}
	var space TransmissionFreeSpace
	if err = c.rpcCall("free-space", payload, &space); err == nil {
		if space.Path == path {
			freeSpace = cunits.ImportInByte(float64(space.Size))
		} else {
			err = fmt.Errorf("returned path '%s' does not match with requested path '%s'", space.Path, path)
		}
	} else {
		err = fmt.Errorf("'free-space' rpc method failed: %v", err)
	}
	return
}

type transmissionFreeSpacePayload struct {
	Path string `json:"path"`
}

// TransmissionFreeSpace represents the freespace available in bytes for a specific path.
// https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L631
type TransmissionFreeSpace struct {
	Path string `json:"path"`
	Size int64  `json:"size-bytes"`
}
