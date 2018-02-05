package main

import (
	"flag"
	"fmt"
)

// system constant
const (
	YaksokVersion = "v.0.1.0" // YaksokVersion is app version
)

const (
	flagVersion = "v" // flag for version
	flagHelp    = "h" // flag for help
)

type YaksokFlagBox struct {
	version *bool
	help    *bool
}

func NewYaksokFlagBox() *YaksokFlagBox {
	box := &YaksokFlagBox{
		version: flag.Bool(flagVersion, false, "yaksok version"),
		help:    flag.Bool(flagHelp, false, "yaksok help"),
	}

	return box
}

type SubFlagBox struct {
	once     *AtFlagSet
	secondly *AtFlagSet
	minutely *AtFlagSet
	hourly   *AtNowFlagSet
	daily    *AtNowFlagSet
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
func (box *FlagBox) Pickup(args []string) error {
	if len(args) > 0 {
		// if argument is subflagset
		return box.subBox.Pickup(args)
	}

	// if argument is main flag withoutsubflags
	return box.yaksokBox.Pickup()
}

//Pickup is Pick flag or flagset from YaksokFlagBox.
func (box *YaksokFlagBox) Pickup() error {
	var err error
	if box.version != nil || box.help != nil {
		if *box.version {
			fmt.Println(YaksokVersion)
		} else if *box.help {
			Usage()
		}
	} else {
		// return fmt.Errorf("")
		panic("Something wrong")
	}

	return err
}

//Pickup is Pick flag or flagset from SubFlagBox.
func (box *SubFlagBox) Pickup(args []string) error {
	var err error

	// Not enought parameters.
	if args == nil {
		err = fmt.Errorf("No arguments")
	} else if len(args) < 2 {
		//에러가 아니라 flasgset 을 보여줘야 합니다
		fmt.Println("Argument is not enough")
		// box.Usage()
	} else {
		var theBox BaseParser

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
			panic("...Who are you?")
		}

		theBox.Parse(args[1:])
	}

	return err
}

//Ready2FlagBox is ready to FlagBox instead of main.
//it makes testable other functions.
func Ready2FlagBox() *FlagBox {
	box := NewFlagBox()

	// overrides flag.Usage to customize yaksok.
	flag.Usage = Usage
	flag.Parse()

	return box
}

func main() {
	box := Ready2FlagBox()
	err := box.Pickup(flag.Args())

	if err != nil {
		fmt.Println("Error:", err.Error())
		Usage()
	}
}
