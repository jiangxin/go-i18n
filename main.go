package main

import (
	"fmt"
	"os"

	"github.com/jiangxin/go-i18n/i18n"
)

func usage(msg string) {
	fmt.Fprintf(os.Stderr, "Usage: go-i18n [<dir>] [<username>]\n")
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", msg)
	os.Exit(1)
}

func main() {
	var (
		i18nDir  string
		domain   = "go-i18n"
		userName string
	)

	switch len(os.Args) {
	case 3:
		userName = os.Args[2]
		fallthrough
	case 2:
		i18nDir = os.Args[1]
	case 1:
		i18nDir = "po/build/locale"
		fmt.Fprintf(os.Stderr, "WARN: use %s as locale root dir\n\n", i18nDir)
	default:
		usage("too many args provided")
	}

	l := i18n.Setup(i18nDir, domain)

	showMessage(l, userName)
}

func showMessage(l *i18n.Locale, userName string) {
	fmt.Println("############################################################")
	fmt.Printf(l.L_("Show messages for lang: %s\n"), l.LocaleName())
	fmt.Println("############################################################")
	fmt.Println("")

	// Translate text from default domain
	fmt.Print(l.L_("Hello, world.\n"))

	// Translate text from default domain
	if userName == "" {
		userName = l.N_("guest")
	}
	fmt.Printf(l.L_("Welcome: %s.\n"), l.L_(userName))

	// Translate text may have plural forms
	for _, n := range []int{1, 2, 3} {
		fmt.Printf(l.Q_("added %d path\n", "added %d paths\n", n), n)
	}

	fmt.Fprintf(os.Stderr, "# DEBUG: locale: %v\n", l.Locale)
	fmt.Println("")
}
