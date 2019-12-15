package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/faabiosr/proto-nats/internal/cli"
	"github.com/faabiosr/proto-nats/internal/gen"
	"github.com/faabiosr/proto-nats/internal/template"
)

const version = "0.0.0"

func main() {
	vFlag := flag.Bool("v", false, "print current version")
	flag.Parse()

	if *vFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	cli := cli.New(
		gen.New(
			version,
			template.New(),
		),
	)

	if err := cli.Run(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
