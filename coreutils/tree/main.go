package main
 
import (
    "fmt"
	"os"
	"flag"
	"path/filepath"
	"strings"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func go_deep(dir_path string, depth int, final int) {
	entries, err := os.ReadDir(dir_path)
	check(err)
	

	for _, entry := range entries {
		if depth != 0 {
			str := strings.Repeat(" -- ", depth)
			str = str + ">"
			fmt.Print(str)
		}
		fmt.Println(entry.Name())
		if entry.IsDir() && depth != final {
			go_deep(filepath.Join(dir_path, entry.Name()), depth + 1, final)
		}
		fmt.Println()
	}
}
func main() {
	var depth int
	flag.IntVar( &depth,"L", 1, "depth")
	flag.Parse()

	var root_path string
	if len(os.Args) == 2 {
		root_path = os.Args[1]
	} else {
		root_path = os.Args[3]
	}
	go_deep(root_path, 0, depth - 1)
}

// to test
//go run coreutils/tree/main.go -L 2 coreutils