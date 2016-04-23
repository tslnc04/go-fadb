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

// LocalFoundFor returns all the LocalFor files that actually exist
func LocalFoundFor(pointer string) (paths []string) {
	for _, l := range LocalFor(pointer) {
		if fileExists(l) {
			paths = append(paths, l)
		}
	}
	return
}

// LocalFileFor returns the first local file found for the pointer
func LocalFileFor(pointer string) string {
	for _, l := range LocalFor(pointer) {
		if fileExists(l) {
			return l
		}
	}
	return ""
}

// LocalFilesFor is plural version of FileFor
func LocalFilesFor(pointers ...string) (files []string) {
	for _, p := range pointers {
		files = append(files, LocalFileFor(p))
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

// AllFoundFor is the same as AllFor but only returns files that exist.
func AllFoundFor(pointer string) (paths []string) {
	for _, l := range AllFor(pointer) {
		if fileExists(l) {
			paths = append(paths, l)
		}
	}
	return
}

// FileFor returns the first of all files that matches a given pointer.
func FileFor(pointer string) string {
	for _, l := range AllFor(pointer) {
		if fileExists(l) {
			return l
		}
	}
	return ""
}

// FilesFor is plural version of FileFor
func FilesFor(pointers ...string) (files []string) {
	for _, p := range pointers {
		files = append(files, FileFor(p))
	}
	return
}
