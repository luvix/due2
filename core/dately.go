/*
Package core
WORNING: package를 만들 때는 반드시 해당 package와 동일한 이름인 소스코드가 존재해야 한다.
*/
package core

import (
	"flag"
)

//TimelyYaksok is a struct for timely flag.
type TimelyYaksok struct {
	name    string   // name of yaksok
	tags    []string // tag of yaksok
	cmd     string   // command of yaksok
	timecmd string   // time command of yaksok
}

func (ys *TimelyYaksok) Name() string {
	return ys.name
}

func (ys *TimelyYaksok) Command() string {
	return ys.cmd
}

func (ys *TimelyYaksok) Tag(tag string) bool {
	for _, t := range ys.tags {
		if tag == t {
			return true
		}
	}
	panic("no tag exist")
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

func NewTimelyFlagSet(command string, errorHandling flag.ErrorHandling) *TimelyFlagSet {
	tfs := new(TimelyFlagSet)
	flagset := flag.NewFlagSet(command, errorHandling)

	tfs.flagset = flagset
	tfs.yaksok = new(TimelyYaksok)

	return tfs
}
