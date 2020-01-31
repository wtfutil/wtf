// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azuredevops

import "strconv"

func negotiateRequestVersion(location *ApiResourceLocation, apiVersion string) (string, error) {
	if apiVersion == "" {
		// if no api-version is sent to the server, the server will decide the version. The server uses the latest
		// released version if the endpoint has been released, otherwise it will use the latest preview version.
		return apiVersion, nil
	}

	matches := apiVersionRegEx.FindStringSubmatch(apiVersion)
	if len(matches) == 0 && matches[0] != "" {
		return apiVersion, &InvalidApiVersion{apiVersion}
	}

	requestedApiVersion, err := NewVersion(matches[1])
	if err != nil {
		return apiVersion, err
	}
	locationMinVersion, err := NewVersion(*location.MinVersion)
	if err != nil {
		return apiVersion, err
	}
	if locationMinVersion.CompareTo(*requestedApiVersion) > 0 {
		// Client is older than the server. The server no longer supports this
		// resource (deprecated).
		return apiVersion, nil
	} else {
		locationMaxVersion, err := NewVersion(*location.MaxVersion)
		if err != nil {
			return apiVersion, err
		}
		if locationMaxVersion.CompareTo(*requestedApiVersion) < 0 {
			// Client is newer than the server. Negotiate down to the latest version
			// on the server
			negotiatedVersion := string(*location.MaxVersion)
			if *location.ReleasedVersion < *location.MaxVersion {
				negotiatedVersion += "-preview"
			}
			return negotiatedVersion, nil
		} else {
			// We can send at the requested api version. Make sure the resource version
			// is not bigger than what the server supports
			negotiatedVersion := matches[1]
			if len(matches) > 3 && matches[3] != "" { // matches '-preview'
				negotiatedVersion += "-preview"
				if len(matches) > 5 && matches[5] != "" { // has a resource version
					requestedResourceVersion, _ := strconv.Atoi(matches[5])
					if *location.ResourceVersion < requestedResourceVersion {
						negotiatedVersion += "." + strconv.Itoa(*location.ResourceVersion)
					} else {
						negotiatedVersion += "." + matches[5]
					}
				}
			} else {
				// requesting released version, ensure server supports a released version, and if not append '-preview'
				locationReleasedVersion, err := NewVersion(*location.ReleasedVersion)
				if err != nil {
					return apiVersion, err
				}
				if (locationReleasedVersion.Major == 0 && locationReleasedVersion.Minor == 0) || locationReleasedVersion.CompareTo(*requestedApiVersion) < 0 {
					negotiatedVersion += "-preview"
				}
			}
			return negotiatedVersion, nil
		}
	}
}
