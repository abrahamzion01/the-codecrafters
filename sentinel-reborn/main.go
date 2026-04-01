package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage go run . input.txt output.txt")
	}
	inputfile := os.Args[1]
	outputfile := os.Args[2]

	content := readFile(inputfile)
	content = binTodecimal(content)

	writeFile(outputfile, content)
}
