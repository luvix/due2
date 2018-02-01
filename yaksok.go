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

//MainFlagBoxFactory makes MainFlagBox new.
func MainFlagBoxFactory() *MainFlagBox {
	return new(MainFlagBox)
}

//Parsable makes timelyflagset parsable.
type Parsable interface {
	Parse(arg []string)
}

// AtFlagSet is for once, minutely, secondly
type AtFlagSet struct {
	flagset *flag.FlagSet
	name    string
	tags    []string
	at      string
}

//NewAtFlagSet makes AtFlagSet new.
func NewAtFlagSet(name string) *AtFlagSet {
	fs := &AtFlagSet{
		name: name,
	}
	fs.flagset = flag.NewFlagSet(name, flag.PanicOnError)
	fs.flagset.Usage = Usage
	return fs
}

func (fs *AtFlagSet) Parse(args []string) {

}

// AtNowFlagSet is for daily, hourly
type AtNowFlagSet struct {
	flagset *flag.FlagSet
	name    string
	tags    []string
	at      string
	now     string
}

//NewAtNowFlagSet makes AtNowFlagSet new.
func NewAtNowFlagSet(name string) *AtNowFlagSet {
	fs := &AtNowFlagSet{
		name: name,
	}
	fs.flagset = flag.NewFlagSet(name, flag.PanicOnError)
	fs.flagset.Usage = Usage
	return fs
}

func (fs *AtNowFlagSet) Parse(args []string) {

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
	fs.flagset.Usage = Usage
	return fs
}

func (fs *AtNowOnFlagSet) Parse(args []string) {

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

//SubFlagBoxFactory makes SubFlagBox new.
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

// FlagBoxFactory makes FlagBox new.
func FlagBoxFactory() *FlagBox {
	box := new(FlagBox)
	box.mainBox = MainFlagBoxFactory()
	box.subBox = SubFlagBoxFactory()

	return box
}

//Pickup is Pick flag or flagset from FlagBox.
func (box *FlagBox) Pickup(args []string) {
	if len(args) > 0 {
		// if argument is subflagset
		box.subBox.Pickup(args)
	} else {
		// if argument is main flag withoutsubflags
		box.mainBox.Pickup()
	}
}

//Pickup is Pick flag or flagset from MainFlagBox.
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

//Pickup is Pick flag or flagset from SubFlagBox.
func (box *SubFlagBox) Pickup(args []string) {
	// Not enought parameters.
	if len(args) < 2 {
		os.Exit(2)
		return
	}

	var theBox Parsable

	switch args[0] {
	case KeyFlagOnce:
		theBox = box.once
	case KeyFlagSecondly:
		theBox = box.secondly
	case KeyFlagMinutely:
		theBox = box.minutely
	case KeyFlagHourly:
		theBox = box.hourly
	case KeyFlagDaily:
		theBox = box.daily
	case KeyFlagWeekly:
		theBox = box.weekly
	case KeyFlagMonthly:
		theBox = box.monthly
	case KeyFlagYearly:
		theBox = box.yearly
	default:
		fmt.Println("엄....")
		return
	}
	theBox.Parse(args[1:])
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
