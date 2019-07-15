package transmissionrpc

import (
	"fmt"

	"github.com/hekmon/cunits"
)

/*
	Session Statistics
	https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L546
*/

// SessionStats returns all (current/cumulative) statistics.
// https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L548
func (c *Client) SessionStats() (stats *SessionStats, err error) {
	if err = c.rpcCall("session-stats", nil, &stats); err != nil {
		err = fmt.Errorf("'session-stats' rpc method failed: %v", err)
	}
	return
}

// SessionStats represents all (current/cumulative) statistics.
// https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L554
type SessionStats struct {
	ActiveTorrentCount int64            `json:"activeTorrentCount"`
	CumulativeStats    *CumulativeStats `json:"cumulative-stats"`
	CurrentStats       *CurrentStats    `json:"current-stats"`
	DownloadSpeed      int64            `json:"downloadSpeed"`
	PausedTorrentCount int64            `json:"pausedTorrentCount"`
	TorrentCount       int64            `json:"torrentCount"`
	UploadSpeed        int64            `json:"uploadSpeed"`
}

// CumulativeStats is subset of SessionStats.
// https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L562
type CumulativeStats struct {
	DownloadedBytes int64 `json:"downloadedBytes"`
	FilesAdded      int64 `json:"filesAdded"`
	SecondsActive   int64 `json:"secondsActive"`
	SessionCount    int64 `json:"sessionCount"`
	UploadedBytes   int64 `json:"uploadedBytes"`
}

// GetDownloaded returns cumulative stats downloaded size in a handy format
func (cs *CumulativeStats) GetDownloaded() (downloaded cunits.Bits) {
	return cunits.ImportInByte(float64(cs.DownloadedBytes))
}

// GetUploaded returns cumulative stats uploaded size in a handy format
func (cs *CumulativeStats) GetUploaded() (uploaded cunits.Bits) {
	return cunits.ImportInByte(float64(cs.UploadedBytes))
}

// CurrentStats is subset of SessionStats.
// https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L570
type CurrentStats struct {
	DownloadedBytes int64 `json:"downloadedBytes"`
	FilesAdded      int64 `json:"filesAdded"`
	SecondsActive   int64 `json:"secondsActive"`
	SessionCount    int64 `json:"sessionCount"`
	UploadedBytes   int64 `json:"uploadedBytes"`
}

// GetDownloaded returns current stats downloaded size in a handy format
func (cs *CurrentStats) GetDownloaded() (downloaded cunits.Bits) {
	return cunits.ImportInByte(float64(cs.DownloadedBytes))
}

// GetUploaded returns current stats uploaded size in a handy format
func (cs *CurrentStats) GetUploaded() (uploaded cunits.Bits) {
	return cunits.ImportInByte(float64(cs.UploadedBytes))
}
