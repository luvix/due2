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

//YaksokFlag is an object model for main level flags
type Yaksok struct {
	version    *bool
	help       *bool
	onece      *flag.FlagSet
	daily      *flag.FlagSet
	weekly     *flag.FlagSet
	monthly    *flag.FlagSet
	yearly     *flag.FlagSet
	secondly   *flag.FlagSet
	minutely   *flag.FlagSet
	hourly     *flag.FlagSet
	list       *flag.FlagSet
	delete     *flag.FlagSet
	preference *flag.FlagSet
}

func YaksokFactory() *Yaksok {
	ys := new(Yaksok)
	ys.secondly = flag.NewFlagSet("secondly", flag.PanicOnError)
	return ys
}

//ShakeFlag is shaking yaksok flag to indicate what yaksok core would be done.
func (flag *Yaksok) ShakeFlag(args []string) {
	if len(args) > 0 {
		// if argument is subflagset
		flag.ShakeSubsetFlag(args)
	} else {
		// if argument is main flag
	}
}

//ShakeSubsetFlag is shaking subset flag of yaksok to indicate what yaksok core would be done.
func (flag *Yaksok) ShakeSubsetFlag(args []string) int {
	switch args[0] {
	case "secondly":
		fmt.Println("sec", args)
		flag.secondly.Parse(args[1:])
		fmt.Println(flag.secondly.Args())
	case "daily":
		flag.daily.Parse(args[1:])
		fmt.Println(flag.daily.Args())
	default:
		fmt.Println("엄....")
	}
	return 1
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
	// secondlyAt := fsSecondly.String("at", time.Now().Format(time.RFC3339), "the time when you want to run a job")
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
