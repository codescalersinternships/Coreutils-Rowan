package main
 
import (
    "fmt"
	"flag"
	"bufio"
	"os"
	"io"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func main() {
 
    var lines int
	flag.IntVar( &lines,"n", 10, "custom number of lines to print")
	flag.Parse()

	var file_name string
	//NOTE HERE: stdin will have to terminate in ctrl D


	var input io.Reader
	var std_lines []string
	var idx int

	if flag.NFlag() == 0 {
		file_name = os.Args[1]
	} else if len(os.Args) == 4 {
		file_name = os.Args[3]
	}
	
	if file_name == "" { 
		input = os.Stdin
		
	} else {
		file, err := os.Open(file_name)
		check(err)
		input = file
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		std_lines = append(std_lines, scanner.Text())
	}
	check(scanner.Err())

	if len(std_lines) < lines {
		idx = 0
	} else {
		idx = len(std_lines) - lines
	}

	for _, line := range std_lines[idx:] {
		fmt.Println(line)
	}
	
}