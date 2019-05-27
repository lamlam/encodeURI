package main

import (
	"flag"
)

func parseArgs() (bool, string) {
	d := flag.Bool("decode", false, "Exec decodeURIComponent")
	flag.Parse()
	return *d, flag.Arg(0)
}
