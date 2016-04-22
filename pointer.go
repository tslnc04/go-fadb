package fadb

import (
	"os"
	fp "path/filepath"
	"strings"
)

func fileExists(path string) bool {
	_, e := os.Stat(path)
	if e != nil {
		return false
	}
	return true
}

func dirExists(path string) bool {
	s, e := os.Stat(path)
	if e != nil || !s.IsDir() {
		return false
	}
	return true
}

// PathFrom resolves an FADB pointer to its actual file system path.
func FileFrom(pointer string) (path string) {
	p := strings.Split(pointer, ".")
	for _, r := range SearchPath {
		if !dirExists(r) {
			continue
		}
		cur := r
		for _, fd := range p {
			cur = fp.Join(cur, fd)
			toml := cur + ".toml"
			json := cur + ".json"
			if fileExists(toml) {
				return toml
			}
			if fileExists(json) {
				return json
			}
			if dirExists(cur) {
				toml := fp.Join(cur, "_.toml")
				json := fp.Join(cur, "_.json")
				if fileExists(toml) {
					return toml
				}
				if fileExists(json) {
					return json
				}
				//TODO check for list files
				//TODO check for true files
				//TODO check for on files
				//TODO check for tsv files
				//TODO check for dsv files
				//TODO check for csv files
				//TODO check for text files beginning
			}
		}
	}
	return ""
}
