package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	// YaksokVersion is app version
	YaksokVersion = "v.0.1.0"

	//KeyFlagOnce is a string
	KeyFlagOnce = "once"
	//KeyFlagSecondly is a string
	KeyFlagSecondly = "secondly"
	//KeyFlagMinutely is a string
	KeyFlagMinutely = "minutely"
	//KeyFlagHourly is a string
	KeyFlagHourly = "hourly"
	//KeyFlagDaily is a string
	KeyFlagDaily = "daily"
	//KeyFlagWeekly is a string
	KeyFlagWeekly = "weekly"
	//KeyFlagMonthly is a string
	KeyFlagMonthly = "Monthly"
	//KeyFlagYearly is a string
	KeyFlagYearly = "yearly"
)

//Usage Overrides flag.Usage.
func Usage() {
	flag.VisitAll(func(f *flag.Flag) {
		// fmt.Fprintf(os.Stderr, "\t%s\t\t%s\n", f.Name, f.Usage)
		fmt.Printf("\t%s\t\t%s\n", f.Name, f.Usage)
	})
}

//YaksokFlagBox is an object model for main level flags
type YaksokFlagBox struct {
	version *bool
	help    *bool
}

//NewYaksokFlagBox makes YaksokFlagBox new.
func NewYaksokFlagBox() *YaksokFlagBox {
	return new(YaksokFlagBox)
}

//Parsable makes timelyflagset parsable.
type Parsable interface {
	Parse(arg []string)
}

//FlagSet is parent struct for AtFlagSet, AtNowFlagSet, AtNowOnFlagSet.
type FlagSet struct {
	Parsable
	flagset *flag.FlagSet
	jobName *string
	jobTags StringArray
}

//Parse parses arguments.
func (fs *FlagSet) Parse(arg []string) {
	fs.flagset.Parse(arg)
}

//Name returns job name.
func (fs *FlagSet) JobName() string {
	return *fs.jobName
}

//Tags returns job tags.
func (fs *FlagSet) JobTags() StringArray {
	return fs.jobTags
}

//NewFlagSet is ...
func NewFlagSet(flagsetname string) *FlagSet {
	fs := &FlagSet{
		flagset: flag.NewFlagSet(flagsetname, flag.PanicOnError),
	}
	fs.jobName = fs.flagset.String("Name", "", "job name")
	fs.flagset.Var(&fs.jobTags, "", "job tag")

	return fs
}

// AtFlagSet is for once, minutely, secondly
type AtFlagSet struct {
	FlagSet
	jobAt *string
}

//NewAtFlagSet makes AtFlagSet new.
func NewAtFlagSet(name string) *AtFlagSet {
	fs := &AtFlagSet{
		FlagSet: *NewFlagSet(name),
	}
	fs.jobAt = fs.flagset.String("at", "", "job schedule")
	fs.flagset.Usage = Usage

	return fs
}

func (fs *AtFlagSet) JobAt() string {
	return *fs.jobAt
}

// AtNowFlagSet is for daily, hourly
type AtNowFlagSet struct {
	AtFlagSet
	jobNow *string
}

//NewAtNowFlagSet makes AtNowFlagSet new.
func NewAtNowFlagSet(name string) *AtNowFlagSet {
	fs := &AtNowFlagSet{
		AtFlagSet: *NewAtFlagSet(name),
	}
	fs.jobNow = fs.flagset.String("now", "", "job schedule")

	return fs
}

func (fs *AtNowFlagSet) JobNow() string {
	return *fs.jobNow
}

// AtNowOnFlagSet is for daily, hourly
type AtNowOnFlagSet struct {
	AtNowFlagSet
	jobOn *string
}

func NewAtNowOnFlagSet(name string) *AtNowOnFlagSet {
	fs := &AtNowOnFlagSet{
		AtNowFlagSet: *NewAtNowFlagSet(name),
	}
	fs.jobOn = fs.flagset.String("on", "", "job schedule")
	return fs
}

func (fs *AtNowOnFlagSet) JobOn() string {
	return *fs.jobOn
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

//NewSubFlagBox makes SubFlagBox new.
func NewSubFlagBox() *SubFlagBox {
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

//FlagBox have YaksokFlagBox and SubFlagBox
type FlagBox struct {
	yaksokBox *YaksokFlagBox
	subBox    *SubFlagBox
}

// NewFlagBox makes FlagBox new.
func NewFlagBox() *FlagBox {
	box := &FlagBox{
		yaksokBox: NewYaksokFlagBox(),
		subBox:    NewSubFlagBox(),
	}

	return box
}

//Pickup is Pick flag or flagset from FlagBox.
func (box *FlagBox) Pickup(args []string) {
	if len(args) > 0 {
		// if argument is subflagset
		box.subBox.Pickup(args)
	} else {
		// if argument is main flag withoutsubflags
		box.yaksokBox.Pickup()
	}
}

//Pickup is Pick flag or flagset from YaksokFlagBox.
func (box *YaksokFlagBox) Pickup() {
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

//Ready2FlagBox is ready to FlagBox instead of main.
//it makes testable other functions.
func Ready2FlagBox() *FlagBox {
	box := NewFlagBox()
	box.yaksokBox.version = flag.Bool("v", false, "Read the module version.")
	box.yaksokBox.help = flag.Bool("h", false, "Gel help")

	// overrides flag.Usage to customize yaksok.
	flag.Usage = Usage
	flag.Parse()

	return box
}

func main() {
	// box := NewFlagBox()
	// box.yaksokBox.version = flag.Bool("v", false, "Read the module version.")
	// box.yaksokBox.help = flag.Bool("h", false, "Gel help")

	// // overrides flag.Usage to customize yaksok.
	// flag.Usage = Usage

	// parse main flag
	// flag.Parse()
	box := Ready2FlagBox()

	box.Pickup(flag.Args())
}
