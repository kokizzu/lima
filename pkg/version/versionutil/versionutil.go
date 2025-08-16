// SPDX-FileCopyrightText: Copyright The Lima Authors
// SPDX-License-Identifier: Apache-2.0

package versionutil

import (
	"strings"

	"github.com/coreos/go-semver/semver"
)

// Parse parses a Lima version string by removing the leading "v" character and
// stripping everything from the first "-" forward (which are `git describe` artifacts and
// not semver pre-release markers) or cutting ".m" suffix from exact version.
// So "v0.19.1-16-gf3dc6ed.m" will be parsed as "0.19.1".
func Parse(version string) (*semver.Version, error) {
	version = strings.TrimPrefix(version, "v")
	version, _, _ = strings.Cut(version, "-")
	version = strings.TrimSuffix(version, ".m")
	return semver.NewVersion(version)
}

func compare(limaVersion, oldVersion string) int {
	if limaVersion == "" {
		if oldVersion == "" {
			return 0
		}
		return -1
	}
	version, err := Parse(limaVersion)
	if err != nil {
		return 1
	}
	// Handle Unparsable oldVersion gracefully - treat as 0.0.0 so Unparsable limaVersion is always greater
	oldVer, err := semver.NewVersion(oldVersion)
	if err != nil {
		return 1
	}
	cmp := version.Compare(*oldVer)
	if cmp == 0 && strings.Contains(limaVersion, "-") {
		cmp = 1
	}
	return cmp
}

// GreaterThan returns true if the Lima version used to create an instance is greater
// than a specific older version. Always returns false if the Lima version is the empty string.
// Unparsable lima versions (like SHA1 commit ids) are treated as the latest version and return true.
// limaVersion is a `github describe` string, not a semantic version. So "0.19.1-16-gf3dc6ed.m"
// will be considered greater than "0.19.1".
func GreaterThan(limaVersion, oldVersion string) bool {
	return compare(limaVersion, oldVersion) > 0
}

// GreaterEqual return true if limaVersion >= oldVersion.
func GreaterEqual(limaVersion, oldVersion string) bool {
	return compare(limaVersion, oldVersion) >= 0
}
