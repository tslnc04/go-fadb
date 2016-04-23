package truefile

import (
    "strings"
    "io"
    "io/ioutil"
)

/* LoadList
 *
 * Reads a string, like an excerpt of a truelist file, one that is just
 * a list of text and returns a map, the key being that element and the
 * value being true. This also removes extra whitespace and comments
 * preceeded with a '#'.
 *
 * Arguments:
 *  d    data; a string of a true file
 *
 * Returns:
 *  map[string]bool    "text": true; map of the text with every value set to true
 */
func LoadList(d string) map[string]bool {
	lines := strings.Split(d, "\n")
	var output map[string]bool = make(map[string]bool)

	for _, line := range lines {
		if !strings.HasPrefix(strings.TrimSpace(line), "#") && len(line) > 0 {
			output[strings.TrimSpace(line)] = true
		}
	}

	return output
}

/* LoadListFile
 *
 * Reads a truelist file, one that is just a list of text and returns a map, 
 * the key being that element and the value being true. This also removes extra 
 * whitespace and comments preceeded with a '#'.
 *
 * Arguments:
 *  p    path; path to true file
 *
 * Returns:
 *  map[string]bool    "text": true; map of the text with every value set to true
 */
func LoadListFile(p string) map[string]bool {
    bdata, err := ioutil.ReadFile(p)
	  if err != nil {panic(err)}

    d := string(bdata)
    lines := strings.Split(d, "\n")
	  var output map[string]bool = make(map[string]bool)

	  for _, line := range lines {
		  if !strings.HasPrefix(strings.TrimSpace(line), "#") && len(line) > 0 {
			  output[strings.TrimSpace(line)] = true
		  }
	  }

	  return output
}

/* LoadListReader
 *
 * Reads an io.Reader of a truelist file, one that is just a list of text
 * and returns a map, the key being that element and the value being true.
 * This also removes extra whitespace and comments preceeded with a '#'.
 *
 * Arguments:
 *  r    reader; io.Reader of file
 *
 * Returns:
 *  map[string]bool    "text": true; map of the text with every value set to true
 */
func LoadListReader(r io.Reader) map[string]bool {
    bdata, err := ioutil.ReadAll(r)
    if err != nil {panic(err)}

    d := string(bdata)
    lines := strings.Split(d, "\n")
	  var output map[string]bool = make(map[string]bool)

	  for _, line := range lines {
		  if !strings.HasPrefix(strings.TrimSpace(line), "#") && len(line) > 0 {
			  output[strings.TrimSpace(line)] = true
		  }
	  }

	  return output
}
