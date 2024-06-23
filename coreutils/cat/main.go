package main
 
import (
    "fmt"
	"bufio"
	"os"
	"io"
	"flag"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func main() {
	var lines int
	flag.IntVar( &lines,"n", -1, "custom number of lines to print")
	flag.Parse()

	var file_name string
	var input io.Reader 

	if len(os.Args) == 1 {
		fmt.Println("No file passed")
		os.Exit(1)
	} else if len(os.Args) == 2 {
		file_name = os.Args[1]
	} else if len(os.Args) == 4 {
		file_name = os.Args[3]
	}

	file, err := os.Open(file_name)
	check(err)
	
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
	
	check(scanner.Err())
	
}