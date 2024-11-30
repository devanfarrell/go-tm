package io

import (
	"bufio"
	"os"
	"strings"
)

// there seems to be better ways to do this
// https://stackoverflow.com/questions/15159118/read-a-character-from-standard-input-in-go-without-pressing-enter
func ReadRuntimeCharacter() string {
	reader := bufio.NewReader(os.Stdin)
	text, error := reader.ReadString('\n')
	if error != nil {
		println("Something went wrong, retrying input")
		return ReadRuntimeCharacter()
	}
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	return text
}
