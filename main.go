package main

import (
	"fmt"
	"os"
)

func help(outputStream *os.File) {
	fmt.Fprintln(outputStream, `gopoggers: celebrate something with the right energy.

	USAGE: gopoggers [ [-a | -b | -p] | [-h | --help] ]

	OPTIONS:
		-h | --help	show this help
		-p		pixel style (default)
		-a		ascii style
		-b		block style`)
}

func main() {
	if len(os.Args) > 2 {
		fmt.Fprintln(os.Stderr, "Error: provide only one argument")
		help(os.Stderr)
		return
	}
	if len(os.Args) == 1 {
		fmt.Println(pixel)
	} else {
		switch os.Args[1] {
		case "-p":
			fmt.Println(pixel)
		case "-a":
			fmt.Println(ascii)
		case "-b":
			fmt.Println(block)
		case "-h":
			help(os.Stdout)
		case "--help":
			help(os.Stdout)
		default:
			fmt.Fprintln(os.Stderr, "Error: unrecognized option")
			help(os.Stderr)
		}
	}

}
