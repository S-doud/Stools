package main

import (
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		return
	}

	toolFunc := toolMap[os.Args[1]]
	toolFunc(os.Args[1:])
}
