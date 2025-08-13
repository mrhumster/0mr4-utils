package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mrhumster/0mr4-utils/tocebab"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: $ tocebab Example string")
		os.Exit(1)
	}

	s := tocebab.Convert(strings.Join(args[1:], "-"))
	fmt.Println(s)
}
