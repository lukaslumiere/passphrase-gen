package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/lukaslumiere/passphrase-gen/pkg/generator"
)

var validSeparators = []string{"-", "_", ".", ",", ":", ";", "|", "+", ""}

var (
	words       int
	separator   string
	capitalize  bool
	noCaps      bool
	specials    int
	count       int
	showVersion bool
	showHelp    bool
)

func init() {
	flag.IntVar(&words, "n", generator.DefaultWords, "")
	flag.IntVar(&words, "num", generator.DefaultWords, "")
	flag.StringVar(&separator, "d", generator.DefaultSep, "")
	flag.StringVar(&separator, "delimiter", generator.DefaultSep, "")
	flag.BoolVar(&capitalize, "c", true, "")
	flag.BoolVar(&capitalize, "caps", true, "")
	flag.BoolVar(&noCaps, "no-caps", false, "")
	flag.IntVar(&specials, "s", 0, "")
	flag.IntVar(&specials, "specials", 0, "")
	flag.IntVar(&count, "count", generator.DefaultCount, "")
	flag.BoolVar(&showVersion, "v", false, "")
	flag.BoolVar(&showVersion, "version", false, "")
	flag.BoolVar(&showHelp, "h", false, "")
	flag.BoolVar(&showHelp, "help", false, "")
	flag.Usage = printUsage
}

func printUsage() {
	fmt.Println("passphrase-gen - Secure passphrase generator")
	fmt.Println()
	fmt.Println("Usage: passphrase-gen [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -n, --num        Number of words (1-20, default: 6)")
	fmt.Println("  -d, --delimiter  Word separator (default: -)")
	fmt.Printf("                   Valid: %s (empty string allowed)\n", strings.Join(validSeparators[:len(validSeparators)-1], " "))
	fmt.Println("  -c, --caps       Capitalize first letter of each word (default: true)")
	fmt.Println("      --no-caps    Disable capitalization")
	fmt.Println("  -s, --specials   Append special characters (0-5, default: 0)")
	fmt.Println("      --count      Number of passphrases (default: 1)")
	fmt.Println("  -v, --version    Show version")
	fmt.Println("  -h, --help       Show help")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  passphrase-gen")
	fmt.Println("  passphrase-gen -n 8 -s 2")
	fmt.Println("  passphrase-gen --count 5 -d \"\"")
	fmt.Println("  passphrase-gen --no-caps -d _")
}

func isValidSeparator(sep string) bool {
	for _, s := range validSeparators {
		if sep == s {
			return true
		}
	}
	return false
}

func run() error {
	flag.Parse()

	if showHelp {
		printUsage()
		return nil
	}

	if showVersion {
		fmt.Println("passphrase-gen version", generator.Version)
		return nil
	}

	if words < generator.MinWords || words > generator.MaxWords {
		return fmt.Errorf("words must be between %d and %d", generator.MinWords, generator.MaxWords)
	}

	if !isValidSeparator(separator) {
		return fmt.Errorf("invalid separator: %q (valid: %s or empty)", separator, strings.Join(validSeparators[:len(validSeparators)-1], " "))
	}

	if specials < 0 || specials > generator.MaxSpecials {
		return fmt.Errorf("specials must be between 0 and %d", generator.MaxSpecials)
	}

	if count < 1 {
		return fmt.Errorf("count must be at least 1")
	}

	caps := capitalize && !noCaps

	cfg := generator.Config{
		Words:      words,
		Separator:  separator,
		Capitalize: caps,
		Specials:   specials,
	}

	passphrases, err := generator.GenerateMultiple(cfg, count)
	if err != nil {
		return err
	}

	for _, p := range passphrases {
		fmt.Println(p)
	}

	entropy := generator.CalculateEntropy(cfg)
	fmt.Printf("\nEntropy: %.1f bits\n", entropy)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
