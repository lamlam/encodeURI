package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	enableDecode, targetURI := parseArgs()

	if len(targetURI) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if enableDecode {
		d, err := decodeURI(targetURI)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(d)
		os.Exit(0)
	}

	e, err := encodeURI(targetURI)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(e)
}
