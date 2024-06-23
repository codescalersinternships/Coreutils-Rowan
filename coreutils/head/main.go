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
	var input io.Reader 
	var std_lines []string

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
	for i := 0; i < lines; i++ {
		if !scanner.Scan() {
			break
		}
		std_lines = append(std_lines, scanner.Text())
	}
	for _, line := range std_lines {
		fmt.Println(line)
	}
	
}