package main

import (
	"fmt"
	"os"
	"renamer/renamer"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("usage:")
		fmt.Println("\trenamer <dirpath> <match_pattern> <replace_pattern>")
		fmt.Println("example")
		fmt.Println("\trenamer '~/' '(.*hoge).txt' '$1_copy.txt'")
		return
	}

	renamer.Rename(args[0], args[1], args[2])
}
