package main

import "os"

type Command struct {
	name      string
	args      []string
	argLength int
}

// function to initialize a Command struct
func command() *Command {
	return &Command{
		name:      os.Args[0],
		args:      os.Args[1:],
		argLength: len(os.Args[1:]),
	}
}
