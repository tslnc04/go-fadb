package dsvparse

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func DSVParse(fp string, delim string) ([]string, [][]string) {
    data, err := ioutil.ReadFile(fp)
	if err != nil {panic(err)}

    sdata := string(data)
    lines := strings.Split(sdata, "\n")
    header := strings.Split(lines[0], delim)
    lines = lines[1:]

    var dsv [][]string

    for _, line := range lines {
        dsv = append(dsv, strings.Split(line, delim))
    }

    return header, dsv[:len(dsv)-1]
}
