package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	//handle no arguments
	if len(args) == 0 {
		fmt.Print("Error: Missing File Name\n\nUsage: SafeStrings <options> <filepath>\n\n")
		os.Exit(0)
	}

	//setup the flags
	minNumPtr := flag.Int("numb", 4, "Set the minimum string length to find")
	flag.Parse()

	filename := args[len(args)-1]

	fileBytes, err := ioutil.ReadFile(filename) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	var buffer []byte
	count := 0
	for _, element := range fileBytes {
		//fmt.Printf("%v: %v\n", index, element)
		if element > 31 && element < 127 {
			buffer = append(buffer, element)
			count = count + 1
		} else {
			if count >= *minNumPtr && element == 0 {
				s := string(buffer[:])
				fmt.Print(s)
				fmt.Print("\n")
				buffer = nil
				count = 0
			} else {
				buffer = nil
				count = 0
			}
		}

	}

	// fmt.Print(fileBytes)
	// fmt.Print("\n")

}
