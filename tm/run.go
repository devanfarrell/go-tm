package tm

import (
	"fmt"
	"go-tm/io"
	"os"
)

func RunTM(definitionString string, inputString string) {

	def, error := ParseDefinition(definitionString)

	if error != nil {
		def.Print()

		println("exiting due to invalid definition")
		return
	}

	exit := false
	for !exit {
		command := io.ReadRuntimeCharacter()

		if command == "d" || command == "D" {
			// tmFacade->deleteStr();
		} else if command == "h" || command == "H" {
			Help()
		} else if command == "i" || command == "I" {
			// tmFacade->insert();
		} else if command == "l" || command == "L" {
			// tmFacade->list();
		} else if command == "q" || command == "Q" {
			// tmFacade->quit();
		} else if command == "r" || command == "R" {
			// tmFacade->run();
		} else if command == "e" || command == "E" {
			// tmFacade->set();
		} else if command == "w" || command == "W" {
			// tmFacade->show();
		} else if command == "t" || command == "T" {
			// tmFacade->truncate();
		} else if command == "v" || command == "V" {
			def.Print()
		} else if command == "x" || command == "X" {
			// tmFacade->exit();
			exit = true
		} else {
			fmt.Fprintf(os.Stderr, "'%s' is not a valid command\n", command)
		}
	}
}
