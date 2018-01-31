package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	// YaksokVersion is app version
	YaksokVersion = "v.0.1.0"

	KeyFlagOnce     = "once"
	KeyFlagSecondly = "secondly"
	KeyFlagMinutely = "minutely"
	KeyFlagHourly   = "hourly"
	KeyFlagDaily    = "daily"
	KeyFlagWeekly   = "weekly"
	KeyFlagMonthly  = "Monthly"
	KeyFlagYearly   = "yearly"
)

//Usage Overrides flag.Usage.
func Usage() {
	flag.VisitAll(func(f *flag.Flag) {
		// fmt.Fprintf(os.Stderr, "\t%s\t\t%s\n", f.Name, f.Usage)
		fmt.Printf("\t%s\t\t%s\n", f.Name, f.Usage)
	})
}

//MainFlagBox is an object model for main level flags
type MainFlagBox struct {
	version *bool
	help    *bool
}

type FlagBox struct {
	mainBox *MainFlagBox
	subBox  *SubFlagBox
}

func MainFlagBoxFactory() *MainFlagBox {
	return new(MainFlagBox)
}

// AtFlagSet is for once, minutely, secondly
type AtFlagSet struct {
	flagset *flag.FlagSet
	name    string
	tags    []string
	at      string
}

func NewAtFlagSet(name string) *AtFlagSet {
	fs := &AtFlagSet{
		name: name,
	}
	fs.flagset = flag.NewFlagSet(name, flag.PanicOnError)
	return fs
}

// AtNowFlagSet is for daily, hourly
type AtNowFlagSet struct {
	flagset *flag.FlagSet
	name    string
	tags    []string
	at      string
	now     string
}

func NewAtNowFlagSet(name string) *AtNowFlagSet {
	fs := &AtNowFlagSet{
		name: name,
	}
	fs.flagset = flag.NewFlagSet(name, flag.PanicOnError)
	return fs
}

// AtNowOnFlagSet is for daily, hourly
type AtNowOnFlagSet struct {
	flagset *flag.FlagSet
	name    string
	tags    []string
	at      string
	now     string
	on      string
}

func NewAtNowOnFlagSet(name string) *AtNowOnFlagSet {
	fs := &AtNowOnFlagSet{
		name: name,
	}
	fs.flagset = flag.NewFlagSet(name, flag.PanicOnError)
	return fs
}

type SubFlagBox struct {
	once     *AtFlagSet
	secondly *AtFlagSet
	minutely *AtFlagSet
	daily    *AtNowFlagSet
	hourly   *AtNowFlagSet
	weekly   *AtNowOnFlagSet
	monthly  *AtNowOnFlagSet
	yearly   *AtNowOnFlagSet
}

func SubFlagBoxFactory() *SubFlagBox {
	box := &SubFlagBox{
		once:     NewAtFlagSet(KeyFlagOnce),
		secondly: NewAtFlagSet(KeyFlagSecondly),
		minutely: NewAtFlagSet(KeyFlagMinutely),
		hourly:   NewAtNowFlagSet(KeyFlagHourly),
		daily:    NewAtNowFlagSet(KeyFlagDaily),
		weekly:   NewAtNowOnFlagSet(KeyFlagWeekly),
		monthly:  NewAtNowOnFlagSet(KeyFlagMonthly),
		yearly:   NewAtNowOnFlagSet(KeyFlagYearly),
	}

	return box
}

// FlagBoxFactory initliazes FlagBox.
func FlagBoxFactory() *FlagBox {
	box := new(FlagBox)
	box.mainBox = MainFlagBoxFactory()
	box.subBox = SubFlagBoxFactory()

	return box
}

//PickupMainFlag is shaking yaksok flag to indicate what yaksok core would be done.
func (box *FlagBox) Pickup(args []string) {
	if len(args) > 0 {
		// if argument is subflagset
		box.subBox.Pickup(args)
	} else {
		// if argument is main flag withoutsubflags
		box.mainBox.Pickup()
	}
}

func (box *MainFlagBox) Pickup() {
	if *box.version {
		fmt.Println(YaksokVersion)
	} else {
		if !*box.help {
			fmt.Fprintf(os.Stderr, "Yaksok needs flag or command.\n")
		}
		Usage()
	}
}

func (box *SubFlagBox) Pickup(args []string) {
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "secondly":
		fmt.Println("sec", args)
		// flag.secondly.Parse(args[1:])
		// fmt.Println(flag.secondly.Args())
	case "daily":
		// flag.daily.Parse(args[1:])
		// fmt.Println(flag.daily.Args())
	default:
		fmt.Println("엄....")
	}
}

func main() {
	box := FlagBoxFactory()
	box.mainBox.version = flag.Bool("v", false, "Read the module version.")
	box.mainBox.help = flag.Bool("h", false, "Gel help")

	// overrides flag.Usage to customize yaksok.
	flag.Usage = Usage

	// parse main flag
	flag.Parse()

	box.Pickup(flag.Args())
}
