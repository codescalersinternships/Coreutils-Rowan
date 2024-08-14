package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var lines int
	flag.IntVar(&lines, "n", -1, "custom number of lines to print")
	flag.Parse()

	var fileName string

	if flag.NFlag() == 0 {
		fileName = os.Args[1]
	} else {
		idx := len(os.Args) - flag.NFlag()
		fileName = os.Args[idx]
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Cannot open file!")
		os.Exit(1)
	}


	scanner := bufio.NewScanner(file)
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
