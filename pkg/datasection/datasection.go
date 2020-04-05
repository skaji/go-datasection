package datasection

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func Parse(data string) map[string]string {
	out := map[string]string{}
	header := regexp.MustCompile(`^@@\s+(.+?)\s*\r?$`)
	var current struct {
		name  string
		lines []string
	}
	trimTail := func(lines []string) []string {
		if len(lines) > 0 && lines[len(lines)-1] == "" {
			return lines[:len(lines)-1]
		}
		return lines
	}
	for _, line := range strings.Split(data, "\n") {
		if header.MatchString(line) {
			if current.name != "" {
				out[current.name] = strings.Join(trimTail(current.lines), "\n") + "\n"
			}
			current.name = header.FindStringSubmatch(line)[1]
			current.lines = nil
		} else {
			current.lines = append(current.lines, line)
		}
	}
	if current.name != "" {
		out[current.name] = strings.Join(trimTail(current.lines), "\n") + "\n"
	}
	return out
}

func ParseFile(path string) (map[string]string, error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return Parse(string(c)), nil
}
