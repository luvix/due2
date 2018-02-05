package main

import (
	"flag"
	"os"
	"testing"
)

func TestTotalCase(t *testing.T) {
	testsets := []struct {
		args []string
	}{
		{[]string{"v"}},
		{[]string{"h"}},
		{[]string{"once"}},
		{[]string{"secondly"}},
		{[]string{"minutely"}},
		{[]string{"hourly"}},
		{[]string{"daily"}},
		{[]string{"weekly"}},
		{[]string{"monthly"}},
		{[]string{"yearly"}},
	}

	box := Ready2FlagBox()
	for _, testSet := range testsets {
		os.Args = testSet.args

		box.Pickup(flag.Args())
	}
}
