package fadb

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mitchellh/go-homedir"
)

//---------------------------- Valid Names ----------------------------

const validNameChars = `[a-zA-Z_][a-zA-Z0-9_-]*`
const protoNameChars = `[A-Z][a-zA-Z0-9_-]*`

var validNameX = regexp.MustCompile(`^` + validNameChars + `$`)
var protoNameX = regexp.MustCompile(`^` + protoNameChars + `$`)

//--------------------------- File Suffixes ---------------------------

var AllFileSuffixes = []string{} // built at init()
var ObjectFileSuffixes = []string{"toml", "json"}
var ListFileSuffixes = []string{"list", "array"}
var TrueFileSuffixes = []string{"true", "on"}
var TableFileSuffixes = []string{"csv", "tsv", "dsv"}
var TextFileSuffixes = []string{
	"md", "mdown", "mdkn", "mkd", "mdwn", "mdtxt", "mdtext",
	"markdown", "mark", "text", "txt", "rst", "html", "asc",
}
var HaveObjectSuffixes = []string{} // built at init()

//---------------------------- Search Path ----------------------------

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

func init() {
	AllFileSuffixes = append(AllFileSuffixes, ObjectFileSuffixes...)
	AllFileSuffixes = append(AllFileSuffixes, ListFileSuffixes...)
	AllFileSuffixes = append(AllFileSuffixes, TableFileSuffixes...)
	AllFileSuffixes = append(AllFileSuffixes, TrueFileSuffixes...)
	AllFileSuffixes = append(AllFileSuffixes, TextFileSuffixes...)
	HaveObjectSuffixes = append(HaveObjectSuffixes, ObjectFileSuffixes...)
	HaveObjectSuffixes = append(HaveObjectSuffixes, TableFileSuffixes...)
	HaveObjectSuffixes = append(HaveObjectSuffixes, TrueFileSuffixes...)
	buildPath()
}
