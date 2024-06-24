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
	flag.IntVar(&lines, "n", -1, "custom number of lines to print")
	flag.Parse()

	var fileName string
	var input io.Reader

	if len(os.Args) == 1 {
		fmt.Println("No file passed")
		os.Exit(1)
	} else if len(os.Args) == 2 {
		fileName = os.Args[1]
	} else if len(os.Args) == 4 {
		fileName = os.Args[3]
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Cannot open file!")
		os.Exit(1)
	}

	input = file

	scanner := bufio.NewScanner(input)
	if lines == -1 {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	} else {
		for i := 0; i < lines; i++ {
			if !scanner.Scan() {
				break
			}
			fmt.Println(scanner.Text())
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Scanner is having some error!")
		os.Exit(1)
	}

}
