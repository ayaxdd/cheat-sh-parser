package main

import (
	"fmt"
	"log"
	"os"
)

const (
	prefix = "https://www.cheat.sh/"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("error: empty parameters; use cheatsh -h")
	}

	config, err := ParseArgs(args)
	if err != nil {
		log.Fatal(err)
	}

	err = config.ValidateArgs()
	if err != nil {
		log.Fatal(err)
	}

	resource, err := GetResource(config.Source)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile("out/out.md", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprint(file, string(resource))
}
