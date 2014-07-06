package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/simonz05/short"
)

var (
	e    = flag.String("e", "", "encode to base62")
	d    = flag.String("d", "", "decode")
	help = flag.Bool("h", false, "show help text")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nOptions:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NFlag() == 0 || *help {
		flag.Usage()
		os.Exit(1)
	}
	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "e":
			n, err := strconv.Atoi(f.Value.String())

			if err != nil {
				fmt.Fprint(os.Stderr, err)
				os.Exit(1)
			}

			fmt.Fprintf(os.Stdout, "%s", short.Encode(n))
		case "d":
			fmt.Fprintf(os.Stdout, "%d", short.Decode([]byte(f.Value.String())))
		}
	})
}
