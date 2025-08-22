package main

import (
	"errors"
	"flag"
)

// Реализация - curl <source> -> format output -> os.Stdin > file name

type UserInput struct {
	Source   string
	FileName string
	isFile   bool
}

func (u *UserInput) ValidateArgs() error {
	if u.Source == "" {
		return errors.New("must specify a name of resourse from cheat.sh/")
	}

	u.isFile = isFlagPassed("o")
	if u.isFile && u.FileName == "" {
		return errors.New("option -o: requires parameter")
	}

	return nil
}

func ParseArgs(args []string) (UserInput, error) {
	u := UserInput{}

	fs := flag.NewFlagSet("cheat-sh", flag.ContinueOnError)
	fs.StringVar(&u.Source, "s", "", "Name of `resourse` to download from cheat.sh/")
	fs.StringVar(&u.FileName, "o", "", "Write to file instead of stdin")

	err := fs.Parse(args)
	if err != nil {
		return u, err
	}

	if fs.NArg() != 0 {
		return u, errors.New("positional arguments specified")
	}

	return u, nil
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})

	return found
}
