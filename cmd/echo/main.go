package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Print(os.Args[i])
		if i == len(os.Args)-1 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
}
