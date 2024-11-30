package main

import (
	"flag"
	"fmt"
	"go-tm/io"
	"go-tm/tm"
	"os"
)

func main() {
	flag.Parse()

	if flag.Arg(0) == "" {
		fmt.Fprintf(os.Stderr, "A file path argument must be passed\n")
	} else {
		tmFileBase := flag.Arg(0)
		tmDefinitionFile := tmFileBase + ".def"
		tmInputFile := tmFileBase + ".str"

		definitionContent, definitionError := io.ReadFile(tmDefinitionFile)
		if definitionError != nil {
			fmt.Fprintf(os.Stderr, "Could not open the file at %s\n", tmDefinitionFile)
			panic(definitionError)
		}

		inputContent, _ := io.ReadFile(tmInputFile)
		tm.RunTM(definitionContent, inputContent)
	}
}
