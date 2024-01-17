package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("Arg %d = %s\n", i, arg)
	}
}