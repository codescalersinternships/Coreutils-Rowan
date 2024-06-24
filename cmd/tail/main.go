package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	var lines int
	flag.IntVar(&lines, "n", 10, "custom number of lines to print")
	flag.Parse()

	var fileName string

	var input io.Reader
	var stdLines []string
	var idx int

	if flag.NFlag() == 0 {
		fileName = os.Args[1]
	} else if len(os.Args) == 4 {
		fileName = os.Args[3]
	}

	if fileName == "" {
		input = os.Stdin

	} else {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Cannot open file!")
			os.Exit(1)
		}
		input = file
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		stdLines = append(stdLines, scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println("Scanner is having some error!")
		os.Exit(1)
	}

	if len(stdLines) < lines {
		idx = 0
	} else {
		idx = len(stdLines) - lines
	}

	for _, line := range stdLines[idx:] {
		fmt.Println(line)
	}

}
