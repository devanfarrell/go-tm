package main

import (
	"flag"
	"fmt"
	"go-tm/io"
	"os"
)

func main() {
	// help := flag.String("help", "help", "The name to greet.")
	// var Usage = func() {
	// 	fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	// 	PrintDefaults()
	// }

	flag.Parse()

	if flag.Arg(0) == "" {
		files, _ := os.Open(".")
		defer files.Close()

		fileInfo, _ := files.Readdir(-1)
		for _, file := range fileInfo {
			fmt.Println(file.Name())
		}
	} else {
		tmFileBase := flag.Arg(0)
		tmDefinitionFile := tmFileBase + ".def"
		tmInputFile := tmFileBase + ".str"
		defExists := io.CheckFileExists(tmDefinitionFile)
		inputExists := io.CheckFileExists(tmInputFile)
		fmt.Fprintf(os.Stderr, "%s, %s, %s, %t, %t\n", flag.Arg(0), tmDefinitionFile, tmInputFile, defExists, inputExists)
		if !defExists {
			fmt.Fprintf(os.Stderr, "Life is pain")
		}
	}
}
