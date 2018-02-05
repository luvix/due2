package main

const (
	KeyFlagList    = "list"
	KeyFlagDelete  = "delete"
	KeyFlagSetting = "setting"
)

type AdminFlagSet struct {
	BaseFlagSet // Base parser for flagset in yaksok
	Gonari      //GoNaRi is a user info.
}

type ListFlagSet struct {
	AdminFlagSet
	Job
	All *bool
}

type DeleteFlagSet struct {
	AdminFlagSet
	Job
}

type SettingFlagSet struct {
}
