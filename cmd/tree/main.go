package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func goDeep(dirPath string, depth int, final int) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Cannot read path!")
		os.Exit(1)
	}

	for _, entry := range entries {
		if depth != 0 {
			str := strings.Repeat(" -- ", depth)
			str = str + ">"
			fmt.Print(str)
		}
		fmt.Println(entry.Name())
		if entry.IsDir() && depth != final {
			goDeep(filepath.Join(dirPath, entry.Name()), depth+1, final)
		}
		fmt.Println()
	}
}
func main() {
	var depth int
	flag.IntVar(&depth, "L", 1, "depth")
	flag.Parse()

	var rootPath string
	if len(os.Args) == 2 {
		rootPath = os.Args[1]
	} else {
		rootPath = os.Args[3]
	}
	goDeep(rootPath, 0, depth-1)
}
