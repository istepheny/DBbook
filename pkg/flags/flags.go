package flags

import (
	"flag"
	"fmt"
	"os"
)

var (
	Host   string
	Port   string
	Ticker int64
	Help   bool
)

func init() {
	flag.StringVar(&Host, "h", "0.0.0.0", "Listening `host`.")

	flag.StringVar(&Port, "p", "3000", "Listening `port`.")

	flag.Int64Var(&Ticker, "t", 3600, "Update documentation every t `seconds`.")
	if Ticker < 1 {
		Ticker = 1
	}

	flag.BoolVar(&Help, "v", false, "Show this help.")

	flag.Usage = usage

	flag.Parse()

	if Help {
		flag.Usage()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr,
		`DBbook version: 1.0.1
Usage: dbbook [-h host] [-p port] [-t seconds] [-v]

Options:
`)
	flag.PrintDefaults()
	os.Exit(0)
}
