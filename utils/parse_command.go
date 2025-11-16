package utils

import "strings"

type ParsedToken struct {
	Command string
	Paths   []string
	Flags   []string
}

func ParseCommand(token string) ParsedToken {
	line := strings.Fields(token)
	if len(line) == 0 {
		return ParsedToken{}
	}
	cmd := line[0]
	paths := []string{}
	flags := []string{}
	for _, item := range line[1:] {
		if strings.HasPrefix(item, "-") {
			flags = append(flags, item)
		} else {
			paths = append(paths, item)
		}
	}
	return ParsedToken{
		Command: cmd,
		Paths:   paths,
		Flags:   flags,
	}
}
