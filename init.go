package fadb

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mitchellh/go-homedir"
)

const (
	validNameChars = `[a-zA-Z_][a-zA-Z0-9_-]*`
	protoNameChars = `[A-Z][a-zA-Z0-9_-]*`
)

var validNameX = regexp.MustCompile(`^` + validNameChars + `$`)
var protoNameX = regexp.MustCompile(`^` + protoNameChars + `$`)

var home, _ = homedir.Dir()

var SearchPath = []string{"."}

func buildPath() {
	env := os.Getenv(`FADBPATH`)
	if env != "" {
		SearchPath = append(SearchPath, strings.Split(env, `:`)...)
	}
	SearchPath = append(SearchPath,
		filepath.Join(home, "fadb"),
		filepath.Join(home, "_fadb"),
		filepath.Join(home, ".fadb"),
		filepath.Join(home, "data"),
		filepath.Join(home, "_data"),
		filepath.Join(home, ".data"),
		"/var/fadb/data",
		"/usr/var/fadb/data",
		"/var/data",
		"/usr/var/data",
		"/usr/share/fadb/data",
		//TODO add the windows data paths
		//TODO add the mac data paths
	)

}

// init is automatically called when fadb is imported
func init() {
	buildPath()
}
