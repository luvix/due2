package core

import (
	"flag"
	"fmt"
)

type TimelyYaksok struct {
	name    string   // name of yaksok
	tags    []string // tag of yaksok
	cmd     string   // command of yaksok
	timecmd *string  // time command of yaksok
}

func (ys *TimelyYaksok) Name() string {
	return ys.name
}

func (ys *TimelyYaksok) Command() string {
	return ys.cmd
}

func (ys *TimelyYaksok) Tag(tag string) string {
	var theTag string
	for _, t := range ys.tags {
		if tag == t {
			theTag = t
		}
	}
	panic("no tag exist")
	return theTag
}

func (ys *TimelyYaksok) Tags() []string {

	tags := make([]string, len(ys.tags))
	copy(tags, ys.tags)

	return tags
}

type TimelyFlagSet struct {
	flagset *flag.FlagSet
	yaksok  *TimelyYaksok
}

//Usage Overrides flag.Usage.
func Usage() {
	flag.VisitAll(func(f *flag.Flag) {
		// fmt.Fprintf(os.Stderr, "\t%s\t\t%s\n", f.Name, f.Usage)
		fmt.Printf("\t%s\t\t%s\n", f.Name, f.Usage)
	})
}

func NewTimelyFlagSet(command string, errorHandling flag.ErrorHandling) *TimelyFlagSet {
	tfs := new(TimelyFlagSet)
	flagset := flag.NewFlagSet(command, errorHandling)

	tfs.flagset = flagset
	tfs.yaksok = new(TimelyYaksok)

	return tfs
}
