// Package version is used as a way to inject the currently built
// version identifier into the compiled binary.
package version

import (
	"fmt"
	"strings"
)

const (
	baseVersion = "1.0.0"
)

var (
	buildNumber = "x"
	buildTime   = ""
)

// Version will return the current compiled version.
func Version() string {
	return strings.TrimSpace(baseVersion)
}

// BuildNumber returns the current build number that the binary has
// been compiled from.
func BuildNumber() string {
	return buildNumber
}

func BuildTime() string {
	return buildTime
}

// FullVersion returns a generated complete version string.
func FullVersion() string {
	return fmt.Sprintf("%s+%s built at %s", Version(), BuildNumber(), BuildTime())
}
