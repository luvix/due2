package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	// YaksokVersion is app version
	YaksokVersion = "v.0.1.0"
)

//Usage Overrides flag.Usage.
func Usage() {
	flag.VisitAll(func(f *flag.Flag) {
		// fmt.Fprintf(os.Stderr, "\t%s\t\t%s\n", f.Name, f.Usage)
		fmt.Printf("\t%s\t\t%s\n", f.Name, f.Usage)
	})
}

func main() {
	fVersion := flag.Bool("v", false, "Read the module version.")
	fHelp := flag.Bool("h", false, "Gel help")

	// fsOnce
	fsDaily := flag.NewFlagSet("daily", flag.ExitOnError)
	// fsWeekly
	// fsMonthly
	// fsYearly
	fsSecondly := flag.NewFlagSet("secondly", flag.ExitOnError)
	// fsMinutely
	// fsHourly
	// fsList
	// fsDelete
	// fsPreference

	// overrides flag.Usage to customize yaksok.
	flag.Usage = Usage

	// parse main flag
	flag.Parse()

	// if argument is exist: user want to use flagset
	if len(flag.Args()) > 0 {
		switch flag.Arg(0) {
		case "secondly":
			fsSecondly.Parse(flag.Args()[1:])
			fmt.Println(fsSecondly.Args())
		case "daily":
			fsDaily.Parse(flag.Args()[1:])
			fmt.Println(fsDaily.Args())
		default:
			fmt.Println("엄....")
		}
	} else {
		if *fVersion {
			fmt.Println(YaksokVersion)
		} else if *fHelp {
			Usage()
		} else {
			fmt.Fprintf(os.Stderr, "Yaksok needs flag or command.\n")
			flag.PrintDefaults()
		}
	}
}
