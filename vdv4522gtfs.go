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
	"patrickbrosi.de/vdv452parser/vdv452"
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

		// for _, l := range feed.Lines {
		// fmt.Fprintf(os.Stdout, "== Line %s (%s)== \n", l.RouteAbbr, l.LineAbbr)
		// for _, s := range l.Sequence {
		// if stop, ok := feed.Stops[uint64(s.PointType)*7000000+uint64(s.PointNo)]; ok {
		// fmt.Fprintf(os.Stdout, "   Point %d | %d (%s)\n", s.PointType, s.PointNo, stop.Stop_Desc)
		// } else {
		// panic(fmt.Errorf("Stop not found: %d | %d", s.PointType, s.PointNo))
		// }
		// }
		// }

		for _, j := range feed.Journeys {
			lineId := fmt.Sprintf("%06d", j.LineNo) + j.RouteAbbr
			journeyDepTime := j.DepartureTime
			l := feed.Lines[lineId]
			if len(l.Sequence) == 0 {
				continue
			}
			var prevStop *vdv452.Stop
			prevTime := 0
			// dest := feed.Destinations[uint64(l.Sequence[0].DestNo)]
			fmt.Fprintf(os.Stdout, "== Trip %d to '%s'== \n", j.JourneyNo)
			for _, s := range l.Sequence {
				if stop, ok := feed.Stops[uint64(s.PointType)*7000000+uint64(s.PointNo)]; ok {
					arrTime := 0
					depTime := 0
					if prevStop == nil {
						arrTime = journeyDepTime
						depTime = arrTime
					} else {

						tGroup := uint64(l.OpDepNo)*1000000000 + uint64(j.TimingGroupNo)
						ft := uint64(prevStop.Point_Type)*100000000000000 + uint64(prevStop.Point_No)*100000000 + uint64(stop.Point_Type)*1000000 + uint64(stop.Point_No)
						travelTime := feed.TravelTimes[tGroup][ft]
						arrTime = prevTime + travelTime
						depTime = arrTime
					}
					fmt.Fprintf(os.Stdout, "   Point %d | %d (%s) %d:%d:%d    %d:%d:%d\n", s.PointType, s.PointNo, stop.Stop_Desc, arrTime/3600, arrTime%3600/60, arrTime%3600%60, depTime/3600, depTime%3600/60, depTime%3600%60)
					prevStop = stop
					prevTime = depTime
				} else {
					panic(fmt.Errorf("Stop not found: %d | %d", s.PointType, s.PointNo))
				}
			}
		}

		fmt.Fprintf(os.Stdout, "done.\n")
	}
}
