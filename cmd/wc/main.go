package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	lPtr := flag.Bool("l", false, "number of lines bool")
	wPtr := flag.Bool("w", false, "number of words bool")
	cPtr := flag.Bool("c", false, "number of characters bool")

	flag.Parse()

	var input io.Reader
	var fileName string
	var l int
	var w int
	var c int

	idx := flag.NFlag()

	if len(os.Args) == idx+1 {
		fmt.Println("No file passed")
		os.Exit(1)
	}

	fileName = os.Args[idx+1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Cannot open file!")
		os.Exit(1)
	}

	input = file

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		l++

		words := strings.Fields(scanner.Text())
		w += len(words)

		c += utf8.RuneCountInString(line)

	}
	if scanner.Err() != nil {
		fmt.Println("Scanner is facing some error!")
		os.Exit(1)
	}
	if *lPtr {
		fmt.Printf("%d ", l)
	}
	if *wPtr {
		fmt.Printf("%d ", w)
	}
	if *cPtr {
		fmt.Printf("%d ", c)
	}
	fmt.Println(fileName)

}
