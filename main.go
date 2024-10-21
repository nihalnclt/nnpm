package main

import "github.com/nihalnclt/nnpm/cmd"

// main is the entry point of the nnpm application.
// It initializes and executes the command-line interface (CLI) defined in the cmd package.
func main() {
	cmd.Execute()
}

// init, install, add, remove, update, list (ls), cache clean, why
