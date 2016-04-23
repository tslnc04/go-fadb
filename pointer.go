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

// LocalFor returns local files that would be searched.
func LocalFor(pointer string) (paths []string) {
	var suf, base string
	p := strings.Split(pointer, ".")

	base = fp.Join(p...)
	for _, suf = range ObjectFileSuffixes {
		paths = append(paths, fp.Join(base, "_."+suf))
	}
	for _, suf = range AllFileSuffixes {
		paths = append(paths, base+"."+suf)
	}

	base = fp.Join(p[:len(p)-1]...)
	for _, suf = range ObjectFileSuffixes {
		paths = append(paths, fp.Join(base, "_."+suf))
	}
	for _, suf = range HaveObjectSuffixes {
		paths = append(paths, base+"."+suf)
	}

	for i := len(p) - 2; i > 0; i-- {
		base = fp.Join(p[:i]...)
		for _, suf = range ObjectFileSuffixes {
			paths = append(paths, fp.Join(base, "_."+suf))
		}
		for _, suf = range ObjectFileSuffixes {
			paths = append(paths, base+"."+suf)
		}
	}
	for _, suf = range ObjectFileSuffixes {
		paths = append(paths, "_."+suf)
	}
	return
}

// AllFor calls LocalFor on everything in SearchPath
func AllFor(pointer string) (paths []string) {
	for _, path := range SearchPath {
		for _, local := range LocalFor(pointer) {
			paths = append(paths, fp.Join(path, local))
		}
	}
	return
}

// FoundFor is the same as AllFor but only returns files that exist.
func FoundFor(pointer string) (paths []string) {
	return
}

// FileFor returns the first file that matches a given pointer.
func FileFor(pointer string) (path string) {
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
