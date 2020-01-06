package main

import (
	"./parserfile"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World from Main")
	myParser := parserfile.Parser{Filename: ""}
	if os.Args[0] != "" {
		myParser.Filename = os.Args[0]
	}
	myParser.Parse()
}
