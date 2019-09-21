// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azuredevops

import (
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
}

func NewVersion(version string) (*Version, error) {
	split := strings.Split(version, ".")
	if len(split) > 1 {
		major, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		minor, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		return &Version{
			Major: major,
			Minor: minor,
		}, nil
	}
	return nil, &InvalidVersionStringError{version: version}
}

func (version Version) CompareTo(compareToVersion Version) int {
	if version.Major > compareToVersion.Major {
		return 1
	} else if version.Major < compareToVersion.Major {
		return -1
	} else if version.Minor > compareToVersion.Minor {
		return 1
	} else if version.Minor < compareToVersion.Minor {
		return -1
	}
	return 0
}

func (version Version) String() string {
	return strconv.Itoa(version.Major) + "." + strconv.Itoa(version.Minor)
}

type InvalidVersionStringError struct {
	version string
}

func (e *InvalidVersionStringError) Error() string {
	return "The version string was invalid: " + e.version
}
