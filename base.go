package main

import (
	"flag"
	"fmt"
)

// [A-Z]tringArray is string array strcut.
// from https://stackoverflow.com/a/28323276
type StringArray []string

func (sa *StringArray) String() string {
	return "StringArray"
}

// [A-Z]et is append a string to string array.
func (sa *StringArray) Set(arg string) error {
	*sa = append(*sa, arg)

	return nil
}

// [A-Z]ob defines job.
type Job struct {
	Name        *string
	Tags        *StringArray
	TimeCommand *StringArray
}

// [A-Z]ewJob creates a new job info.
func NewJob() *Job {
	job := &Job{
		Name:        new(string),
		Tags:        new(StringArray),
		TimeCommand: new(StringArray),
	}
	return job
}

// Gonari is a user who administrate whose or those job.
type Gonari struct {
	who   *string
	those *string
}

// NewGonari creates a new gonari info.
func NewGonari(name string, group string) *Gonari {
	gnr := &Gonari{
		who:   &name,
		those: &group,
	}
	return gnr
}

// Usage Overrides flag.Usage.
func Usage() {
	flag.VisitAll(func(f *flag.Flag) {
		// fmt.Fprintf(os.Stderr, "\t%s\t\t%s\n", f.Name, f.Usage)
		fmt.Printf("\t%s\t\t%s\n", f.Name, f.Usage)
	})
}

// BaseParser supports every has-a flag.FlagSet struct.
type BaseParser interface {
	Parse(args []string) error
}

// BaseFlagSet has common member variables(flag.FlagSet) and methods(BaseParser.Parse).
type BaseFlagSet struct {
	BaseParser               // Base parser for flagset in yaksok
	flagset    *flag.FlagSet // flagset for yaksok
}

// [A-Z]arse is deprecated.
func (fs *BaseFlagSet) Parse(args []string) error {
	return fs.flagset.Parse(args)
}

// NewBaseFlagSet supports child FlagSets to allocate new one.
// it allocates with flag.FlagSet which has name and flag.PanicOnError for ErrorHandling.
func NewBaseFlagSet(name string) *BaseFlagSet {
	fs := &BaseFlagSet{
		flagset: flag.NewFlagSet(name, flag.PanicOnError),
	}

	return fs
}
