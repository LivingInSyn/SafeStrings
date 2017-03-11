package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	filename := args[0]

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
			if count > 4 && element == 0 {
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
