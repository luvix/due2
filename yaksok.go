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

//Yaksok is an object model for main level flags
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

type SubFlagBox struct {
	//TODO: 곧 만들꺼야
	flagsets map[string]*flag.FlagSet
}

func (box *SubFlagBox) Names() []string {
	return []string{
		"once", "daily", "weekly", "monthly", "yearly",
		"hourly", "minutely", "secondly",
		"list", "delete", "setting",
	}
}

func (box *SubFlagBox) UpdateFlagSets(name string) {
	box.flagsets[name] = flag.NewFlagSet(name, flag.PanicOnError)
}

func SubFlagBoxFactory() *SubFlagBox {
	box := new(SubFlagBox)
	box.flagsets = make(map[string]*flag.FlagSet)

	for _, v := range box.Names() {
		box.UpdateFlagSets(v)
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
