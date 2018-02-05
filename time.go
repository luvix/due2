package main

// these constants are key for flag.FlagSet
const (
	KeyFlagOnce     = "once"     //string fro flag.FlagSet once
	KeyFlagSecondly = "secondly" //string fro flag.FlagSet secondly
	KeyFlagMinutely = "minutely" //string fro flag.FlagSet minutely
	KeyFlagHourly   = "hourly"   //string fro flag.FlagSet hourly
	KeyFlagDaily    = "daily"    //string fro flag.FlagSet daily
	KeyFlagWeekly   = "weekly"   //string fro flag.FlagSet weekly
	KeyFlagMonthly  = "monthly"  //string fro flag.FlagSet monthly
	KeyFlagYearly   = "yearly"   //string fro flag.FlagSet yaerly
)

type TimeFlagSet struct {
	BaseFlagSet // Base parser for flagset in yaksok
	Job         // job is job.
}

func NewTimeFlagSet(name string) *TimeFlagSet {
	fs := &TimeFlagSet{
		BaseFlagSet: *NewBaseFlagSet(name),
		Job:         *new(Job),
	}
	return fs
}

type AtFlagSet struct {
	TimeFlagSet
	At string
}

func NewAtFlagSet(name string) *AtFlagSet {
	fs := &AtFlagSet{
		TimeFlagSet: *NewTimeFlagSet(name),
	}
	return fs
}

type AtNowFlagSet struct {
	AtFlagSet
	Now string
}

func NewAtNowFlagSet(name string) *AtNowFlagSet {
	fs := &AtNowFlagSet{
		AtFlagSet: *NewAtFlagSet(name),
	}
	return fs
}

type AtNowOnFlagSet struct {
	AtNowFlagSet
	On string
}

func NewAtNowOnFlagSet(name string) *AtNowOnFlagSet {
	fs := &AtNowOnFlagSet{
		AtNowFlagSet: *NewAtNowFlagSet(name),
	}
	return fs
}
