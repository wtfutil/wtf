//go:build !windows

package git

import "os"

const (
	__path_seperator = string(os.PathSeparator)
	__go_cmd         = "git"
)
