package main
 
import (
    "fmt"
	"flag"
	"os"
	"bufio"
	"io"
	"strings"
	"unicode/utf8"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func main() {
 

	lines_ptr := flag.Bool("l", false, "number of lines bool")
	words_ptr := flag.Bool("w", false, "number of words bool")
	chars_ptr := flag.Bool("c", false, "number of characters bool")

	flag.Parse()


	var input io.Reader 
	var file_name string
	var l int
	var w int
	var c int

	idx := flag.NFlag()

	if len(os.Args) == idx + 1 {
		fmt.Println("No file passed")
		os.Exit(1)
	}

	file_name = os.Args[idx + 1]

	file, err := os.Open(file_name)
	check(err)

	input = file

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		l ++

		words := strings.Fields(scanner.Text())
		w += len(words)

		c += utf8.RuneCountInString(line)

	}
	check(scanner.Err())
	if *lines_ptr {
		fmt.Printf("%d ",l)
	}
	if *words_ptr {
		fmt.Printf("%d ",w)
	}
	if *chars_ptr {
		fmt.Printf("%d ",c)
	}
	fmt.Println(file_name)
	
}