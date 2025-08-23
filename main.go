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

	cfg, err := NewConfig(args)
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.ValidateArgs()
	if err != nil {
		log.Fatal(err)
	}

	resource, err := GetResource(cfg.Source)
	if err != nil {
		log.Fatal(err)
	}

	w := os.Stdin
	if cfg.IsFile == true {
		w, err = os.OpenFile(cfg.FileName, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer w.Close()
	}

	fmt.Fprint(w, string(resource))
}
