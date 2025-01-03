package definition

import (
	"errors"
	"strings"
)

func ParseDefinition(input string) (Definition, error) {
	description := parseDescription(input)
	states, _ := parseStates(input)

	definition := Definition{description: description, states: states}
	return definition, errors.New("Not yet implemented")
}

func parseDescription(input string) string {
	description := ""
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if nextSectionStart(line) {
			return description
		}
		description = description + "\n" + line
	}
	return description
}

func parseStates(input string) ([]string, error) {
	states := []string{}
	lines := strings.Split(input, "\n")

	statesFound := false
	for _, line := range lines {
		lowerString := strings.ToLower(line)
		if strings.HasPrefix(lowerString, "states:") {
			statesFound = true

			internalStrings := strings.Split(line, " ")
			for i, state := range internalStrings {
				if i != 0 && state != "" {
					states = append(states, state)
				}
			}
		} else if statesFound && !nextSectionStart(line) {
			internalStrings := strings.Split(line, " ")
			for _, state := range internalStrings {
				if state != "" {
					states = append(states, state)
				}
			}

		} else if statesFound && nextSectionStart(line) {
			return states, nil
		}
	}

	if len(states) == 0 {
		return states, errors.New("No definition for states found")
	}
	return states, nil
}

func nextSectionStart(line string) bool {
	lowerString := strings.ToLower(line)
	return strings.HasPrefix(lowerString, "tape_alphabet:") || strings.HasPrefix(lowerString, "transition_function:") || strings.HasPrefix(lowerString, "inital_state:") || strings.HasPrefix(lowerString, "blank_character:") || strings.HasPrefix(lowerString, "final_states:") || strings.HasPrefix(lowerString, "states:") || strings.HasPrefix(lowerString, "input_alphabet:")
}
