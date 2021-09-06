package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("You input: ")
	if len(os.Args) < 2 {
		fmt.Println("Missing input file parameter.")
		return
	}
	data, error := ioutil.ReadFile(os.Args[1])
	if error != nil {
		fmt.Println("Can't read file: ", os.Args[1])
		panic(error)
	}
	set := make(map[string]bool)

	fmt.Println("Reading strings...")
	var codes = strings.Split(string(data), "\n")
	for _, item := range codes {
		set[item] = true
	}

	var code = string("{\n")

	keys := make([]string, 0, len(set))
	for i := range set {
		keys = append(keys, i)
	}

	// for java array specifically
	for _, key := range keys {
		code += "\t\"" + key + "\",\n"
	}
	code += "}"
	fmt.Println("Succesfully read strings, writing to output.txt...")

	output, err := os.Create("output.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer output.Close()

	_, writeErr := output.WriteString(code)

	if writeErr != nil {
		fmt.Println(writeErr)
	}

	fmt.Println("Done.")
}
