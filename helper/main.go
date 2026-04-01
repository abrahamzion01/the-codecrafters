// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {

// 	if len(os.Args) != 3 {
// 		fmt.Println("go run . input.txt output.txt")
// 	} else {
// 		fmt.Println("Converted successfully")
// 	}

// 	inputfile := os.Args[1]
// 	outputfile := os.Args[2]

// 	content := readFile(inputfile)
// 	content = fixAToAn(content)

// 	writeFile(outputfile, content)
// }
