package main

import (
	"flag"
	"fmt"
	// "os"
)

// StringArray is string array strcut.
// from https://stackoverflow.com/a/28323276
type StringArray []string

func (sa *StringArray) String() string {
	return "StringArray"
}

// Set is append a string to string array.
func (sa *StringArray) Set(arg string) error {
	*sa = append(*sa, arg)

	return nil
}

// Job defines job.
type Job struct {
	Name        *string
	Tags        *StringArray
	TimeCommand *StringArray
}

// NewJob creates a new job info.
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

// BaseParser supports every has-a flag.FlagSet struct.
type BaseParser interface {
	Parse(args []string) error
	FlagSetName() string
	String() string
	Usage()
}

// BaseFlagSet has common member variables(flag.FlagSet) and methods(BaseParser.Parse).
type BaseFlagSet struct {
	BaseParser               // Base parser for flagset in yaksok
	flagset    *flag.FlagSet // flagset for yaksok
	name       string        // flagset name
}

// Parse is deprecated.
func (fs *BaseFlagSet) Parse(args []string) error {
	defer func() error {
		if r := recover(); r != nil {
		}
		return nil
	}()

	err := fs.flagset.Parse(args)
	return err
}

func (fs *BaseFlagSet) FlagSetName() string {
	return fs.name
}

func DefaultUsage(fs *flag.FlagSet) {
	fs.VisitAll(func(f *flag.Flag) {
		fmt.Printf("-%s\t%s\n", f.Name, f.Usage)
	})
}

func (fs *BaseFlagSet) Usage() {
	fmt.Println("flags in yaksok:", fs.name)
	DefaultUsage(fs.flagset)
}

func (fs *BaseFlagSet) String() string {
	var options string

	fs.flagset.VisitAll(func(f *flag.Flag) {
		options += " -" + f.Name + " "
		options += f.Value.String()
	})

	stringified := fs.name + options
	return stringified
}

// NewBaseFlagSet supports child FlagSets to allocate new one.
// it allocates with flag.FlagSet which has name and flag.PanicOnError for ErrorHandling.
func NewBaseFlagSet(name string) *BaseFlagSet {
	fs := &BaseFlagSet{
		flagset: flag.NewFlagSet(name, flag.PanicOnError),
	}
	fs.name = name
	fs.flagset.Usage = fs.Usage
	return fs
}
