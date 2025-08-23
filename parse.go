package main

import (
	"errors"
	"flag"
)

// Реализация - curl <source> -> format output -> os.Stdin > file name

type Config struct {
	Flags    *flag.FlagSet
	Source   string
	FileName string
	IsFile   bool
}

func NewConfig(args []string) (Config, error) {
	c := Config{}

	fs := flag.NewFlagSet("cheat-sh", flag.ContinueOnError)
	fs.StringVar(&c.Source, "s", "", "Name of `resourse` to download from cheat.sh/")
	fs.StringVar(&c.FileName, "o", "", "Write to `file` instead of stdin")

	err := fs.Parse(args)
	if err != nil {
		return c, err
	}

	if fs.NArg() != 0 {
		return c, errors.New("positional arguments specified")
	}

	c.Flags = fs

	return c, nil
}

func (c *Config) ValidateArgs() error {
	if c.Source == "" {
		return errors.New("must specify a name of resourse from cheat.sh/")
	}

	passed := c.isFlagPassed("o")
	if passed && c.FileName == "" {
		return errors.New("option -o: requires parameter")
	}
	c.IsFile = passed && c.FileName != ""

	return nil
}

func (c *Config) isFlagPassed(name string) bool {
	found := false
	c.Flags.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})

	return found
}
