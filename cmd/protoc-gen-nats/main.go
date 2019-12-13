package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.0.0"

func main() {
	vFlag := flag.Bool("v", false, "print current version")
	flag.Parse()

	if *vFlag {
		fmt.Println(version)
		os.Exit(0)
	}
}
