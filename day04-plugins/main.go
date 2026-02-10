package main

import (
	"fmt"
	"os"

	processor "example.com/day04-plugins/processor"
)

func selectProcessor(name string) processor.Processor {
	switch name {
	case "upper":
		return processor.Uppercase{}
	case "lower":
		return processor.Lowercase{}
	case "rot13":
		return processor.Rot13{}
	default:
		return nil
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: program <module> <text>")
		return
	}

	module := os.Args[1]
	text := os.Args[2]

	processor := selectProcessor(module)
	if processor == nil {
		fmt.Println("Unknown module:", module)
		return
	}

	result := processor.Process(text)
	fmt.Println(result)
}
