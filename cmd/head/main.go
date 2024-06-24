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
	for i := 0; i < lines; i++ {
		if !scanner.Scan() {
			break
		}
		stdLines = append(stdLines, scanner.Text())
	}
	for _, line := range stdLines {
		fmt.Println(line)
	}

}
