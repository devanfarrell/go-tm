package tm

import (
	"errors"
	"strings"
)

func ParseDefinition(input string) (Definition, error) {
	description := parseDescription(input)

	definition := Definition{description: description}
	return definition, errors.New("Not yet implemented")
}

func parseDescription(input string) string {
	description := ""
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "STATES:") {
			return description
		}
		description = description + "\n" + line
	}
	return description
}
