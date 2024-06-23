package main
 
import (
    "fmt"
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
	var file_name string
	var input io.Reader 

	if len(os.Args) == 1 {
		fmt.Println("No file passed")
		os.Exit(1)
	}
	file_name = os.Args[1]
	file, err := os.Open(file_name)
	check(err)
	
	input = file

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	check(scanner.Err())
	
}