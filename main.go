package main

import (
	"fmt"
	"os"
)

func usage(msg string) {
	fmt.Fprintf(os.Stderr, "Usage: go-i18n [<dir>] [<username>]\n")
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", msg)
	os.Exit(1)
}

func main() {
	var (
		userName string
	)

	switch len(os.Args) {
	case 2:
		userName = os.Args[1]
	case 1:
	default:
		usage("too many args provided")
	}

	showMessage(userName)
}

func showMessage(userName string) {
	fmt.Println("############################################################")
	fmt.Print("Show messages\n")
	fmt.Println("############################################################")
	fmt.Println("")

	// Translate text from default domain
	fmt.Print("Hello, world.\n")

	// Translate text from default domain
	if userName == "" {
		userName = "guest"
	}
	fmt.Printf("Welcome: %s.\n", userName)

	// Translate text may have plural forms
	for _, n := range []int{1, 2, 3} {
		if n == 1 {
			fmt.Printf("added %d path\n", n)
		} else {
			fmt.Printf("added %d paths\n", n)
		}
	}

	fmt.Println("")
}
