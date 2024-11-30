package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	name := flag.String("name", "World", "Name to greet")
	age := flag.Int("age", 0, "Your age")
	help := flag.Bool("help", false, "Show Help")

	// Parse the command-line flags
	flag.Parse()

	// Use the parsed flags
	if *help {
		fmt.Printf("Help: %v", *help)
	} else {
		fmt.Printf("Hello, %s!\n", *name)
		fmt.Printf("You are %d years old.\n", *age)
	}
}
