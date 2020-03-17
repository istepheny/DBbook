package flags

import (
	"flag"
	"fmt"
	"os"
)

var (
	Help   bool
	Port   string
	Ticker int64
)

func init() {
	flag.StringVar(&Port, "p", "3000", "Listening `port`.")

	flag.Int64Var(&Ticker, "t", 3600, "Update documentation every t `seconds`.")
	if Ticker < 1 {
		Ticker = 1
	}

	flag.BoolVar(&Help, "h", false, "Show this help.")

	flag.Usage = usage

	flag.Parse()

	if Help {
		flag.Usage()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr,
		`DBbook version: 1.0.0
Usage: dbbook [-h] [-p port] [-t seconds]

Options:
`)
	flag.PrintDefaults()
	os.Exit(0)
}
