package main

import (
	"flag"
)

func parseArgs() (bool, string) {
	d := flag.Bool("d", false, "Exec decodeURI")
	flag.Parse()
	return *d, flag.Arg(0)
}
