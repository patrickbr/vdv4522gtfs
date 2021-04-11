// Copyright 2016 Patrick Brosi
// Authors: info@patrickbrosi.de

//
// Use of this source code is governed by a GPL v2
// license that can be found in the LICENSE file

package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
	"patrickbrosi.de/vdv452parser"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "vdv4522gtfs - (C) 2016-2021 by P. Brosi <info@patrickbrosi.de>\n\nUsage:\n\n  %s [<options>] <input VDV452>\n\nAllowed options:\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	help := flag.BoolP("help", "?", false, "this message")

	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	vdv452paths := flag.Args()

	if len(vdv452paths) == 0 {
		fmt.Fprintln(os.Stderr, "No VDV452 location specified, see --help")
		os.Exit(1)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, "Error:", r)
		}
	}()

	for _, path := range vdv452paths {
		fmt.Fprintf(os.Stdout, "Parsing VDV452 in '%s' ...", vdv452paths)
		feed := vdv452parser.NewVDV452()
		feed.Parse(path)

		fmt.Fprintf(os.Stdout, "done.\n")
	}
}
